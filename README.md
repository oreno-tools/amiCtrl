# amiCtrl
[![CircleCI](https://circleci.com/gh/oreno-tools/amiCtrl.svg?style=svg)](https://circleci.com/gh/oreno-tools/amiCtrl)
## これなに

- AMI の作成、削除、詳細を確認出来るワンバイナリツールです

## 使い方

### インストール

https://github.com/inokappa/amiCtrl/releases から環境に応じたバイナリをダウンロードしてください.

```
wget https://github.com/inokappa/amiCtrl/releases/download/v0.0.5/amiCtrl_darwin_amd64 ~/bin/amiCtrl
chmod +x ~/bin/amiCtrl
```

### ヘルプ

```sh
$ ./amiCtrl -h
Usage of ./amiCtrl:
  -ami string
        AMI ID を指定.
  -batch
        バッチモードで実行.
  -create
        AMI を作成.
  -days int
        日数を指定. (要: --prefix オプションと併用)
  -delete
        AMI を削除.
  -endpoint string
        AWS API のエンドポイントを指定.
  -instance string
        Instance ID を指定.
  -json
        JSON 形式で出力.
  -latest
        最新の AMI を取得 (要: --prefix オプションと併用)
  -name string
        AMI Name を指定.
  -noreboot
        No Reboot オプションを指定. (default true)
  -prefix string
        AMI Name の Prefix を指定.
  -profile string
        Profile 名を指定.
  -region string
        Region 名を指定. (default "ap-northeast-1")
  -sort-by-creation
        作成日順にソートして出力.
  -version
        バージョンを出力.
```

### AMI 作成

```sh
$ ./amiCtrl -instance=i-18173987 -name=suzuki-ami-desu -create
+-----------------+--------------+-----------+-------------+
|    AMI NAME     |    AMI ID    |   STATE   | SNAPSHOT ID |
+-----------------+--------------+-----------+-------------+
| suzuki-ami-desu | ami-00385acd | available | snap-0c1cf6 |
+-----------------+--------------+-----------+-------------+
```

### AMI 情報取得

```sh
$ ./amiCtrl -ami=ami-1234567x
+-----------------+--------------+-----------+-------------+
|    AMI NAME     |    AMI ID    |   STATE   | SNAPSHOT ID |
+-----------------+--------------+-----------+-------------+
| suzuki-ami-desu | ami-00385acd | available | snap-0c1cf6 |
+-----------------+--------------+-----------+-------------+
```

`-ami` オプションを指定しない場合, 自アカウントに登録されている AMI 全てを出力します.

また, `-json` をオプションを付与すると, JSON 形式で AMI の情報を取得します.

```sh
$ ./amiCtrl -ami=ami-1234567x -json | jq .
{
  "amis": [
    {
      "ami_name": "suzuki-ami",
      "ami_id": "ami-00385acd",
      "instance_type": "available",
      "snapshot_ids": [
        "snap-07865e9992ce1b6cf"
      ]
    }
  ]
}
```


### AMI 削除

```sh
$ ./amiCtrl -ami=ami-1234567x -delete
+-----------------+--------------+-----------+-------------+
|    AMI NAME     |    AMI ID    |   STATE   | SNAPSHOT ID |
+-----------------+--------------+-----------+-------------+
| suzuki-ami-desu | ami-00385acd | available | snap-0c1cf6 |
+-----------------+--------------+-----------+-------------+
上記の AMI を削除しますか?(y/n): y
AMI を削除します...
AMI を削除しました.
```
