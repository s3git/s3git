s3git: git for Cloud Storage
============================

s3git applies the git philosophy to Cloud Storage. If you know git, you will know how to use s3git!

s3git is a simple CLI tool that allows you to create a *distributed*, *decentralized* and *versioned* repository. It scales limitlessly to 100s of millions of files and PBs of storage and stores your data safely in S3. Yet huge repos can be cloned on the SSD of your laptop for making local changes, committing and pushing back.

Exactly like git, s3git does not require any server-side components, just download and run the executable. It imports the golang package [s3git-go](https://github.com/s3git/s3git-go) that can be used from other applications as well.

Download binaries
-----------------

**DISCLAIMER: These are PRE-RELEASE binaries -- use at your own peril for now**

### OSX

Download `s3git` from [https://github.com/s3git/s3git/releases/download/v0.9.0/s3git-darwin-amd64](https://github.com/s3git/s3git/releases/download/v0.9.0/s3git-darwin-amd64)

```sh
$ wget -q -O s3git https://github.com/s3git/s3git/releases/download/v0.9.0/s3git-darwin-amd64
$ chmod +x s3git
$ ./s3git
```

### Linux

Download `s3git` from [https://github.com/s3git/s3git/releases/download/v0.9.0/s3git-linux-amd64](https://github.com/s3git/s3git/releases/download/v0.9.0/s3git-linux-amd64)

```sh
$ wget -q -O s3git https://github.com/s3git/s3git/releases/download/v0.9.0/s3git-linux-amd64
$ chmod +x s3git
$ ./s3git
```

Building from source
--------------------

Unfortunately not yet possible due to missing SDK, stay tuned!

Example workflow
----------------

Here is a simple workflow to create a new repository and populate it with some data:
```sh
$ s3git init
Initialized empty s3git repository in ...
$ # Just stream in some text
$ echo "hello s3git" | s3git add
Added: 18e622875a89cede0d7019b2c8afecf8928c21eac18ec51e38a8e6b829b82c3ef306dec34227929fa77b1c7c329b3d4e50ed9e72dc4dc885be0932d3f28d7053
$ # Add some more files
$ s3git add "*.mp4"
$ # Commit and log
$ s3git commit -m "My first commit"
$ s3git log --pretty
```

Push to cloud storage
---------------------

```sh
$ # Add remote back end and push to it
$ s3git remote add "primary" -r s3://s3git-playground -a "AKIAJYNT4FCBFWDQPERQ" -s "OVcWH7ZREUGhZJJAqMq4GVaKDKGW6XyKl80qYvkW"
$ s3git push
$ # Read back content
$ s3git cat 18e6
hello s3git
```

_Note: Do not store any important info in the s3git-playground bucket. It will be auto-deleted within 24-hours._
 
Clone the YFCC100M dataset
--------------------------

Clone a large repo with 100 million files totaling 11.5 TB in size ([Multimedia Commons](http://aws.amazon.com/public-data-sets/multimedia-commons/)), yet requiring only 7 GB local disk space (takes several minutes):

```sh
$ s3git clone s3://s3git-100m -a "AKIAI26TSIF6JIMMDSPQ" -s "5NvshAhI0KMz5Gbqkp7WNqXYlnjBjkf9IaJD75x7"
Cloning into ...
Done. Totaling 97,974,749 objects.
$ cd s3git-100m
$ # List all files starting with '123456'
$ s3git ls 123456
12345649755b9f489df2470838a76c9df1d4ee85e864b15cf328441bd12fdfc23d5b95f8abffb9406f4cdf05306b082d3773f0f05090766272e2e8c8b8df5997
123456629a711c83c28dc63f0bc77ca597c695a19e498334a68e4236db18df84a2cdd964180ab2fcf04cbacd0f26eb345e09e6f9c6957a8fb069d558cadf287e
123456675eaecb4a2984f2849d3b8c53e55dd76102a2093cbca3e61668a3dd4e8f148a32c41235ab01e70003d4262ead484d9158803a1f8d74e6acad37a7a296
123456e6c21c054744742d482960353f586e16d33384f7c42373b908f7a7bd08b18768d429e01a0070fadc2c037ef83eef27453fc96d1625e704dd62931be2d1
$ s3git cat cafebad > olympic.jpg
$ # List and count total nr of files
$ s3git ls | wc -l
97974749
```

Fork that repo
--------------

Below is an example for `alice` and `bob` working together on a repository.

```sh
$ mkdir alice && cd alice
alice $ s3git clone s3://s3git-spoon-knife -a "AKIAJYNT4FCBFWDQPERQ" -s "OVcWH7ZREUGhZJJAqMq4GVaKDKGW6XyKl80qYvkW"
Cloning into .../alice/s3git-spoon-knife
Done. Totaling 0 objects.
alice $ cd s3git-spoon-knife
alice $ # add a file filled with zeros
alice $ dd if=/dev/zero count=1 | s3git add
Added: 3ad6df690177a56092cb1ac7e9690dcabcac23cf10fee594030c7075ccd9c5e38adbaf58103cf573b156d114452b94aa79b980d9413331e22a8c95aa6fb60f4e
alice $ # add 9 more files (with random content)
alice $ for n in {1..9}; do dd if=/dev/urandom count=1 | s3git add; done
alice $ # commit
alice $ s3git commit -m "Commit from alice"
alice $ # and push
alice $ s3git push
```

Clone it again as `bob` on a different computer/different directory/different universe:
 
```sh
$ mkdir bob && cd bob
bob $ s3git clone s3://s3git-spoon-knife -a "AKIAJYNT4FCBFWDQPERQ" -s "OVcWH7ZREUGhZJJAqMq4GVaKDKGW6XyKl80qYvkW"
Cloning into .../bob/s3git-spoon-knife
Done. Totaling 10 objects.
bob $ cd s3git-spoon-knife
bob $ # Check if we can access our empty file
bob $ s3git cat 3ad6 | hexdump
00000000  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00
*
00000200
bob $ # add another 10 files
bob $ for n in {1..10}; do dd if=/dev/urandom count=1 | s3git add; done
bob $ # commit
bob $ s3git commit -m "Commit from bob"
bob $ # and push back
bob $ s3git push
```

Swtich back to `alice` again to pull the new content:

```sh
alice $ s3git pull
Done. Totaling 20 objects.
alice $ s3git log --pretty
3f67a4789e2a820546745c6fa40307aa490b7167f7de770f118900a28e6afe8d3c3ec8d170a19977cf415d6b6c5acb78d7595c825b39f7c8b20b471a84cfbee0 Commit from bob
a48cf36af2211e350ec2b05c98e9e3e63439acd1e9e01a8cb2b46e0e0d65f1625239bd1f89ab33771c485f3e6f1d67f119566523a1034e06adc89408a74c4bb3 Commit from alice
```

_Note: Do not store any important info in the s3git-spoon-knife bucket. It will be auto-deleted within 24-hours._

Happy forking!

<a name="minio"></a> Integration with Minio
-------------------------------------------

Instead of S3 you can happily use the [Minio](https://minio.io) server, for example the public server at https://play.minio.io:9002. Just make sure you have a bucket created using [mc](https://github.com/minio/mc) (example below uses `s3git-test`):

```sh
$ s3git init 
$ s3git remote add "primary" -r s3://s3git-test -a "Q3AM3UQ867SPQQA43P2F" -s "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG" -e "https://play.minio.io:9002"
$ echo "hello minio" | s3git add
Added: c7bb516db796df8dcc824aec05db911031ab3ac1e5ff847838065eeeb52d4410b4d57f8df2e55d14af0b7b1d28362de1176cd51892d7cbcaaefb2cd3f616342f
$ s3git commit -m "Commit for minio test"
[commit 6eb708ec7dfd75d9d6a063e2febf16bab3c7a163e203fc677c8a9178889bac012d6b3fcda56b1eb160b1be7fa56eb08985422ed879f220d42a0e6ec80c5735ea]
$ s3git push
Pushing 1 / 1 [==============================================================================================================================] 100.00 % 0
```

and clone it 

```sh
$ s3git clone s3://s3git-test -a "Q3AM3UQ867SPQQA43P2F" -s "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG" -e "https://play.minio.io:9002"
Cloning into .../s3git-test
Done. Totaling 1 object.
$ cd s3git-test/
$ s3git ls
c7bb516db796df8dcc824aec05db911031ab3ac1e5ff847838065eeeb52d4410b4d57f8df2e55d14af0b7b1d28362de1176cd51892d7cbcaaefb2cd3f616342f
$ s3git cat c7bb
hello minio
$ s3git log --pretty
6eb708ec7dfd75d9d6a063e2febf16bab3c7a163e203fc677c8a9178889bac012d6b3fcda56b1eb160b1be7fa56eb08985422ed879f220d42a0e6ec80c5735ea Commit for minio test
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

**Q** Do you support encryption?  
**A** Simply encrypt it before streaming into `s3git add`, eg pipe it through `openssl enc` or similar.  
**Q** Do you support zipping?  
**A** Simply zip it before streaming into `s3git add`, eg pipe it through `zip -r - .` or similar.  
**Q** Why don't you provide a FUSE interface?  
**A** Supporting FUSE would mean introducing a lot of complexity related to POSIX which we would rather avoid.  
