package main

import (
    "os"
    "fmt"
    "flag"
    "strings"
    "encoding/json"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/service/ec2"

    "github.com/olekukonko/tablewriter"
)

const AppVersion = "0.0.5"

var (
    argProfile = flag.String("profile", "", "Profile 名を指定.")
    argRegion = flag.String("region", "ap-northeast-1", "Region 名を指定.")
    argEndpoint = flag.String("endpoint", "", "AWS API のエンドポイントを指定.")
    argInstance = flag.String("instance", "", "Instance ID を指定.")
    argAmi = flag.String("ami", "", "AMI ID を指定.")
    argName = flag.String("name", "", "AMI Name を指定.")
    argCreate = flag.Bool("create", false, "AMI を作成.")
    argDelete = flag.Bool("delete", false, "AMI を削除.")
    argNoreboot = flag.Bool("noreboot", true, "No Reboot オプションを指定.")
    argVersion = flag.Bool("version", false, "バージョンを出力.")
    argJson = flag.Bool("json", false, "JSON 形式で出力する.")
    argBatch = flag.Bool("batch", false, "バッチモードで実行する.")
)

type Amis struct {
    Amis           []Ami    `json:"amis"`
}

type Ami struct {
    AmiName        string `json:"ami_name"`
    AmiId          string `json:"ami_id"`
    AmiState       string `json:"instance_type"`
    SnapshotIds    []string `json:"snapshot_ids"`
}

func awsEc2Client(profile string, region string) *ec2.EC2 {
    var config aws.Config
    if profile != "" {
        creds := credentials.NewSharedCredentials("", profile)
        config = aws.Config{Region: aws.String(region),
                            Credentials: creds,
                            Endpoint: aws.String(*argEndpoint)}
    } else {
        config = aws.Config{Region: aws.String(region), Endpoint: aws.String(*argEndpoint)}
    }
    sess := session.New(&config)
    ec2Client := ec2.New(sess)
    return ec2Client
}

func outputTbl(data [][]string) {
    table := tablewriter.NewWriter(os.Stdout)
    table.SetHeader([]string{"AMI Name", "AMI ID", "State", "Snapshot ID"})
    table.SetAutoMergeCells(true)
    table.SetRowLine(true)

    for _, value := range data {
        table.Append(value)
    }
    table.Render()
}

func outputJson(data [][]string) {
    var rs []Ami
    for _, record := range data {
        snapshot_ids := strings.Split(record[3], "\n")
        r := Ami{AmiName:record[0],
                 AmiId:record[1],
                 AmiState:record[2],
                 SnapshotIds:snapshot_ids}
        rs = append(rs, r)
    }
    rj := Amis{
        Amis: rs,
    }
    b, err := json.Marshal(rj)
    if err != nil {
        fmt.Println("JSON Marshal error:", err)
        return
    }
    os.Stdout.Write(b)
}

func displayAmiInfo(Amis [][]string) {
    if *argJson {
        outputJson(Amis)
    } else {
        outputTbl(Amis)
    }
}

func createTag(ec2Client *ec2.EC2, amiId string, name string) {
    var amiIds []*string
    var tags []*ec2.Tag
    var tag *ec2.Tag

    tag = &ec2.Tag{
        Key: aws.String("Name"),
        Value: aws.String(name),
    }
    tags = append(tags, tag)
    amiIds = append(amiIds, aws.String(amiId))

    input := &ec2.CreateTagsInput{
        Resources: amiIds,
        Tags: tags,
    }
    _, err := ec2Client.CreateTags(input)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
    // fmt.Println(result)
}

func describeAmi(ec2Client *ec2.EC2, amiId string) {
    input := &ec2.DescribeImagesInput{
        Owners: []*string{
            aws.String("self"),
        },
    }

    if amiId != "" {
        input.SetFilters(
            []*ec2.Filter{
                {
                    Name: aws.String("image-id"),
                    Values: []*string{
                        aws.String(amiId),
                    },
                },
            },
        )
    }

    result, err := ec2Client.DescribeImages(input)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    allAmis := [][]string{}
    for _, i := range result.Images {
        var snapshotIds []string
        amiId = *i.ImageId
        snapshotIds = getSnapshotIds(ec2Client, amiId)
        Ami := []string{
            *i.Name,
            amiId,
            *i.State,
            strings.Join(snapshotIds, "\n"),
        }
        allAmis = append(allAmis, Ami)
    }

    displayAmiInfo(allAmis)
}

func createAmi(ec2Client *ec2.EC2, instanceId string, name string, noReboot bool) {
    input := &ec2.CreateImageInput{
        InstanceId: aws.String(instanceId),
        Name: aws.String(name),
        Description: aws.String("Created by amiCtrl."),
        NoReboot: aws.Bool(noReboot),
    }
    res, err := ec2Client.CreateImage(input)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
    createTag(ec2Client, *res.ImageId, name)
    describeAmi(ec2Client, *res.ImageId)
}

func deleteAmi(ec2Client *ec2.EC2, amiId string) {
    var snapshotIds []string
    snapshotIds = getSnapshotIds(ec2Client, amiId)
    if ! *argBatch {
        describeAmi(ec2Client, amiId)
        fmt.Print("上記の AMI を削除しますか?(y/n): ")
        var stdin string
        fmt.Scan(&stdin)
        switch stdin {
        case "y", "Y":
            fmt.Println("AMI を削除します...")
            input := &ec2.DeregisterImageInput{
                ImageId: aws.String(amiId),
            }
            _, err := ec2Client.DeregisterImage(input)
            if err != nil {
                fmt.Println(err.Error())
                os.Exit(1)
            }
            deleteSnapshot(ec2Client, snapshotIds)
            fmt.Println("AMI を削除しました.")
        case "n", "N":
            fmt.Println("処理を停止します.")
            os.Exit(0)
        default:
            fmt.Println("処理を停止します.")
            os.Exit(0)
        }
    } else {
        input := &ec2.DeregisterImageInput{
            ImageId: aws.String(amiId),
        }
        _, err := ec2Client.DeregisterImage(input)
        if err != nil {
            fmt.Println(err.Error())
            os.Exit(1)
        }
        deleteSnapshot(ec2Client, snapshotIds)
        fmt.Println("AMI を削除しました.")
    }
}

func getSnapshotIds(ec2Client *ec2.EC2, amiId string) (snapshotIds []string) {
    description := "Created by CreateImage(*) for " + amiId + " from *"
    input := &ec2.DescribeSnapshotsInput{
        Filters: []*ec2.Filter{
            {
                Name: aws.String("description"),
                Values: []*string{
                    aws.String(description),
                },
            },
        },
    }
    result, err := ec2Client.DescribeSnapshots(input)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    for _, s := range result.Snapshots {
        snapshotIds = append(snapshotIds, *s.SnapshotId)
    }

    return snapshotIds
}

func deleteSnapshot(ec2Client *ec2.EC2, snapshotIds []string) {
    for _, SnapshotId := range snapshotIds {
        input := &ec2.DeleteSnapshotInput{
            SnapshotId: aws.String(SnapshotId),
        }
        _, err := ec2Client.DeleteSnapshot(input)
        if err != nil {
            fmt.Println(err.Error())
            os.Exit(1)
        }
    }
}

func main() {
    flag.Parse()

    if *argVersion {
        fmt.Println(AppVersion)
        os.Exit(0)
    }

    ec2Client := awsEc2Client(*argProfile, *argRegion)

    if *argCreate {
        if *argInstance == "" {
           fmt.Println("`-instance` オプションを指定して下さい.")
            os.Exit(1)
        }
        createAmi(ec2Client, *argInstance, *argName, *argNoreboot)
    } else if *argDelete {
        if *argAmi == "" {
           fmt.Println("`-ami` オプションを指定して下さい.")
            os.Exit(1)
        }
        deleteAmi(ec2Client, *argAmi)
    } else {
        describeAmi(ec2Client, *argAmi)
        os.Exit(0)
    }
}
