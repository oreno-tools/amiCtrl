INSTANCE_ID=$(aws --profile=dummy_profile --region=us-east-1 --endpoint=http://192.168.0.100:5000 \
    ec2 run-instances \
      --image-id=ami-1a2b3c4d \
      --count=1 \
      --instance-type=c3.large \
      --key-name=MyKeyPair \
      --security-groups=MySecurityGroup \
      --query=Instances[].InstanceId --output=text)

AMI_ID=$(aws --profile=dummy_profile --region=us-east-1 --endpoint=http://192.168.0.100:5000 \
  ec2 create-image \
    --instance-id=$(echo ${INSTANCE_ID} | tr -d \\r) \
    --name=test-image --output=text)

gom run amiCtrl.go -profile=dummy_profile -region=us-east-1 -endpoint=http://192.168.0.100:5000 \
  -delete