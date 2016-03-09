s3git: git for Cloud Storage
============================

s3git applies the git philosophy to Cloud Storage. If you know git, you will know how to use s3git!

s3git is a simple CLI tool that allows you to create a *distributed*, *decentralized* and *versioned* repository. It scales limitlessly to 100s of millions of files and PBs of storage and stores your data safely in S3. Yet huge repos can be cloned on the SSD of your laptop for making local changes, committing and pushing back.

Exactly like git, s3git does not require any server-side components, just download and run the executable. It imports the golang package [s3git-go](https://github.com/s3git/s3git-go) that can be used from other applications as well.

Download binaries
-----------------
<<describe how to download

Building from source
--------------------

Make sure you have a working Golang environment and do as follows
```
go get -d github.com/s3git/s3git
cd $GOPATH/src/github.com/s3git/s3git
go install
```

Example workflow
----------------

Here is a simple workflow to create a new repository and populate it with some data:
```
$ s3git init s3://mybucket -a "ACCESSKEY" -s "SECRETKEY"
$ s3git add "*.jpg"
$ s3git commit -m "My first commit"
$ s3git push
$ s3git log
```

Clone the YFCC100M dataset
--------------------------

Clone a large repo with 100 million files totaling 11.5 TB in size ([Multimedia Commons](http://aws.amazon.com/public-data-sets/multimedia-commons/)), yet requiring only 7 GB local disk space (takes several minutes):

```
$ s3git clone s3://s3git-100m-usw2
Cloning into ...
Done. Totalling 97345456 objects.
$ s3git ls 123456
|100 kB| 12345649755b9f489df2470838a76c9df1d4ee85e864b15cf328441bd12fdfc23d5b95f8abffb9406f4cdf05306b082d3773f0f05090766272e2e8c8b8df5997
|100 kB| 123456629a711c83c28dc63f0bc77ca597c695a19e498334a68e4236db18df84a2cdd964180ab2fcf04cbacd0f26eb345e09e6f9c6957a8fb069d558cadf287e
|100 kB| 123456675eaecb4a2984f2849d3b8c53e55dd76102a2093cbca3e61668a3dd4e8f148a32c41235ab01e70003d4262ead484d9158803a1f8d74e6acad37a7a296
|100 kB| 123456e6c21c054744742d482960353f586e16d33384f7c42373b908f7a7bd08b18768d429e01a0070fadc2c037ef83eef27453fc96d1625e704dd62931be2d1
$ s3git cat cafebad > olympic.jpg
```

And collaborate
---------------

Continuing as `alice` from the example above, clone it again as `bob` on a different computer or in a different directory

```
alice $
```

```
bob $
```

Contributions
-------------

Contributions are welcome! Please see [`CONTRIBUTING.md`](CONTRIBUTING.md).

Key features
------------

 * **Easy:** Use a workflow and syntax that you already know and love

 * **Fast:** Lightning fast operation, especially on large files and huge repositories

 * **Infinite scalability:** Stop worrying about maximum repository sizes and have the ability to grow indefinitely

 * **Work from local SSD:** Make a huge cloud disk appear like a local drive

 * **Instant sync:** Push local changes and pull down instantly on other clones

 * **Versioning:** Keep previous versions safe and have the ability to undo or go back in time

 * **Forking:** Ability to make many variants by forking

 * **Verifiable:** Be sure that you have everything and be tamper-proof (“data has not been messed with”)

 * **Deduplication:** Do not store the same data twice

 * **Simplicity:** Simple by design and provide one way to accomplish tasks

Command Line Help
-----------------

```
$ s3git help
s3git applies the git philosophy to Cloud Storage. If you know git, you will know how to use s3git.

s3git is a simple CLI tool that allows you to create a distributed, decentralized and versioned repository.
It scales limitlessly to 100s of millions of files and PBs of storage and stores your data safely in S3.
Yet huge repos can be cloned on the SSD of your laptop for making local changes, committing and pushing back.

Usage:
  s3git [command]

Available Commands:
  add         Add file(s) to the repository
  cat         Read a file from the repository
  clone       Clone a repository into a new directory
  commit      Commit the changes in the repository
  init        Create an empty repository
  ls          List files in the repository
  pull        Update local repository
  push        Update remote repositories
  remote      Manage remote reposities
  status      Show changes to repository

Flags:
  -h, --help[=false]: help for s3git

Use "s3git [command] --help" for more information about a command.
```

Use cases
---------

```
s3git commit "Holiday pictures"
s3git commit "Photos from birthday"
s3git log
```

License
-------

s3git is released under the Apache License v2.0. You can find the complete text in the file LICENSE.

FAQ
---

**Q** Why don't you provide a FUSE interface?  
**A** Supporting FUSE would mean introducing a lot of complexity related to POSIX which we would rather avoid.
