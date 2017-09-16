# amiCtrl

## これなに

- AMI の作成、削除、詳細を確認出来るワンバイナリツールです

## 使い方

### 環境に応じて

Linux 版(linux ディレクトリ以下)、MacOSX 版(osx ディレクトリ以下)、Windows 版(win ディレクトリ)の各バイナリを任意のパスに展開する。

```sh
$ tree -L 2
.
├── Gomfile
├── README.md
├── amiCtrl.go
├── build.sh
├── linux
│   └── amiCtrl
├── osx
│   └── amiCtrl
└── win
    └── amiCtrl.exe

3 directories, 7 files
```

### ヘルプ

```sh
$ ./amiCtrl -h
Usage of ./amiCtrl:
  -ami string
        AMI ID を指定.
  -create
        タグをインスタンスに付与.
  -delete
        タグをインスタンスから削除.
  -describe
        タグを詳細を確認.
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
```

### AMI 作成

```sh
$ ./amiCtrl -instance=i-18173987 -name=suzuki-ami-desu -create
+-----------------+--------------+--------------------------------+
|    AMI NAME     |    AMI ID    |          SNAPSHOT ID           |
+-----------------+--------------+--------------------------------+
| suzuki-ami-desu | ami-1234567x | snap-123456789a1234567         |
|                 |              | snap-123456789b1234567         |
+-----------------+--------------+--------------------------------+
```

### AMI 情報取得

```sh
$ ./amiCtrl -ami=ami-1234567x
+-----------------+--------------+--------------------------------+
|    AMI NAME     |    AMI ID    |          SNAPSHOT ID           |
+-----------------+--------------+--------------------------------+
| suzuki-ami-desu | ami-1234567x | snap-123456789a1234567         |
|                 |              | snap-123456789b1234567         |
+-----------------+--------------+--------------------------------+
```

### AMI 削除

```sh
$ ./amiCtrl -ami=ami-1234567x -delete
+-----------------+--------------+--------------------------------+
|    AMI NAME     |    AMI ID    |          SNAPSHOT ID           |
+-----------------+--------------+--------------------------------+
| suzuki-ami-desu | ami-1234567x | snap-123456789a1234567         |
|                 |              | snap-123456789b1234567         |
+-----------------+--------------+--------------------------------+
上記の AMI を削除しますか?(y/n): y
AMI を削除します...
AMI を削除しました.
```
