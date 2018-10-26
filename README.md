# amiCtrl
[![Build Status](https://travis-ci.org/inokappa/amiCtrl.svg?branch=master)](https://travis-ci.org/inokappa/amiCtrl)
## これなに

- AMI の作成、削除、詳細を確認出来るワンバイナリツールです

## 使い方

### インストール

https://github.com/inokappa/amiCtrl/releases から環境に応じたバイナリをダウンロードしてください.

```
wget https://github.com/inokappa/amiCtrl/releases/download/v0.0.3/amiCtrl_darwin_amd64 ~/bin/amiCtrl
chmod +x ~/bin/amiCtrl
```

### ヘルプ

```sh
$ ./amiCtrl -h
Usage of ./amiCtrl:
  -ami string
        AMI ID を指定.
  -create
        AMI を作成.
  -delete
        AMI を削除.
  -endpoint string
        AWS API のエンドポイントを指定.
  -instance string
        Instance ID を指定.
  -name string
        AMI Name を指定.
  -noreboot
        No Reboot オプションを指定. (default true)
  -profile string
        Profile 名を指定.
  -region string
        Region 名を指定. (default "ap-northeast-1")
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
