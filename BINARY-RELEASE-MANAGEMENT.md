
Binary Release Management
=========================

You can use s3git to manage multiple versions of builds of a particular software product. As an example here we have taken the releases from the [Kubernetes](https://github.com/kubernetes/kubernetes) project.

Add a new release
-----------------

Here is how to add a new release and push it into the cloud (assuming you have cloned earlier, see below).

```sh
$ cd s3git-kubernetes
$ # Make sure we have an empty directory
$ rm -rf *
$ # Download latest release
$ curl -sL https://github.com/kubernetes/kubernetes/releases/download/v1.2.4/kubernetes.tar.gz | tar xz -C ..
$ # Create snaphost
$ s3git snapshot create -m "kubernetes v1.2.4" .
$ # And push to cloud
$ s3git push
```

Pull and checkout latest release
--------------------------------

If you want to pull the latest release and check it out, do as follows (assuming you have cloned earlier, see below)

```sh
$ cd s3git-kubernetes
$ s3git pull
$ s3git snapshot checkout .
```

All releases
------------

If you want to clone the repo on the new machine, do this

```sh
$ s3git clone s3://s3git-kubernetes -a "AKIAJYNT4FCBFWDQPERQ" -s "OVcWH7ZREUGhZJJAqMq4GVaKDKGW6XyKl80qYvkW"
Cloning into .../s3git-kubernetes
Done. Totaling 2,773 objects.
$ cd s3git-kubernetes
$ du -sh .
396K .
$ s3git log --pretty
5882e0fa550c2997c93830a35b12cbe1185c8f64c1e35ccffe29b97ef26d81ef4dfbb75cd5dfca8993430a82a2f48ff0b7 kubernetes v1.2.4
f68893ebf30d1aa43ca70eaaacde7092dc25640f3d8b43b2d4f2edb552723f7c693558a3ded8bbfe4779b0eca221f93e29 kubernetes v1.2.3
8330ba28541c0662d11f4c61357c78f601a44b633ba6a09257f0ea8fc442b0e41b18bcac4b34a875e4161d5bd921dc6468 kubernetes v1.2.2
12c564ac90917a20c01e31be7043485497b0f854aa97bdf6c3d931342133ccdf0d49f6b7ebd01cca2b7ce50f1de59a3359 kubernetes v1.2.1
de39be9f14e674ca749a48a284627a79fdb7807b715d71efaf366205b3ad99890b59022cdbbbd2c7e6c9ece8ba5ce1f158 kubernetes v1.2.0
bd4dcc628b6c4cf5e3ce179d92c39a3cc81c5b30df8c12b9dbd892ecf9e2bea46baf8980cbfe13a617ab5882bd8a2f5d92 kubernetes v1.1.8
663b9e0274cddc36dc1888c0a9ba35f94b890d94d3bffe607cab40005f42501939270e53fb7b0985e7a0aae3346cc65171 kubernetes v1.1.7
b64033494c0cb6feb167121ca048332bb3059c994bf9ddcb474a1ccda450bf0e4bcd528b4581634dbad619713d563c5944 kubernetes v1.1.4
a8e65be1a1eb8b1a4640d03a9431333e99884f4173146357a179c69e4057801a322200a1b1e73217408e29a331d94e40e1 kubernetes v1.1.3
e5c2add2c101e71429d1f7a8fa58ddc63a428f87d74fa7ed99527c1d9037e8673f8835d3bb207b3eb3f87d8c874564f94d kubernetes v1.1.2
7cb200f5e30dbd800a41cfd6ac3c345127da3dfe000ae7854c5b0bc937ff3517c85b478ecde26affaa5d6b51326ed38715 kubernetes v1.1.1
79335bfd40b1a3e36457bb0d2739e35ba1b3ac2fdd503823cc9d335d1808ddf2eaf3fa65d29ad0559196dbaccb5705e8b8 kubernetes v1.0.7
4fcc0a895112a5087664113c54d8fc79cb6d04100e6ccc7ebcb149e45e7a9d46bcf0256dbe6b9516d5a42c4e861451587d kubernetes v1.0.6
92ad92a8c9ef75134b15b3fa59b4f0134f34ab61b6fcb4dd65d027ed6375cb4ddd5c0b0f063b70c8126cb2548960597674 kubernetes v1.0.3
7b1fcf3fc3d07d501ad2ca75205a2f4bf7067a0c2267ebbcb516697bad36005fe716c8b239ce664e0a6ee1adf36130f328 kubernetes v1.0.1
$ s3git snapshot checkout . 
$ platforms/darwin/amd64/kubectl version -c
Client Version: version.Info{Major:"1", Minor:"2", GitVersion:"v1.2.4", GitCommit:"3eed1e3be6848b877ff80a93da3785d9034d0a4f", GitTreeState:"clean"}
$ s3git snapshot checkout . HEAD^
$ platforms/darwin/amd64/kubectl version -c
Client Version: version.Info{Major:"1", Minor:"2", GitVersion:"v1.2.3", GitCommit:"882d296a99218da8f6b2a340eb0e81c69e66ecc7", GitTreeState:"clean"}
$ s3git snapshot checkout . 7b1fcf3f
$ platforms/darwin/amd64/kubectl version -c
Client Version: version.Info{Major:"1", Minor:"0", GitVersion:"v1.0.1", GitCommit:"6a5c06e3d1eb27a6310a09270e4a5fb1afa93e74", GitTreeState:"clean"}
```

Diff a file between revisions
-----------------------------

If you know the hashes of two files you would like to diff use `s3git cat` to fetch the files and pipe them to diff (in this case they correspond to the version file in the root directory):

```sh
$ diff <(s3git cat de749f48) <(s3git cat a3eeab16)
< v1.2.4
---
> v1.2.3
```

Using the `s3git snaphost ls --hash` you can list the files of a snapshot (together with the corresponding hash) and `grep` the file that you are looking, get its hash and pass it on to `s3git cat`.

```sh
$ diff <(s3git cat `s3git snapshot ls --hash HEAD | grep version | grep -v docs | awk '{print $2}'`) <(s3git cat `s3git snapshot ls --hash HEAD^ | grep version | grep -v docs | awk '{print $2}'`)
< v1.2.4
---
> v1.2.3
```

Grab straight out of Cloud Storage
----------------------------------

If you push with the `--hydrate` option you can actually grab the files straight out of S3/Cloud Storage. In order for this you need to use the `--presigned` option for the `s3git snapshot ls` command. Then a simply `wget` or `curl` will then allow you to fetch the object.

```sh
$ # Look for README file  (NB need to exclude docs, examples and cluster subdirs as these contain READMEs as well)
$ s3git snapshot ls --presigned HEAD | grep README | grep -v 'docs\|examples\|cluster'
/users/test/kubernetes/README.md --> https://s3git-kubernetes.s3.amazonaws.com/dc8f38d3866bc8958d13fd6f35ffbb117dcafe6670a49a6df2ea8b1d20df069bf89e6184925516d46ca009623d9c03238e6d7699c41c310fda7f9c1fa84ecd1c?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAJYNT4FCBFWDQPERQ%2F20160526%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20160526T082654Z&X-Amz-Expires=3600&X-Amz-SignedHeaders=host&X-Amz-Signature=fe43befb292a1e570b99a8633756b6746113d609e64cd59de14e3426f7e41a00
$
$ # Grab URL (third element) and pass to wget and dump to stdout
$ s3git snapshot ls --presigned HEAD | grep README | grep -v 'docs\|examples\|cluster' | awk '{print $3}' | wget -i - -O - -q
```

(Note that the `--hydrate` option comes at the expense of losing deduplication savings.)

Deduped format for updating a snapshot
--------------------------------------

By default a `s3git snapshot checkout` will reconstruct the directory structure and files into their original format (and download the content from cloud storage).

If all you are interested in is write access (ie. writing/updating/renaming/removing files) then you can use the `--dedupe` flag for the `s3git snapshot checkout`. This will recreate the directory and files structure with shadow content that contains binary hashes/pointers, so if you (recursively) list the snapshot it will appear identical to a full-fledged checkout (note that many files are 128 bytes or larger multiples of 64 bytes).

You can now copy in new data and or move files around or delete them. When you are done you can create a new snapshot as regular using `s3git snapshot create -m "Updated version"` in order to push it upstream.

