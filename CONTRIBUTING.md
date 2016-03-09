
Setup your Github Repository
----------------------------

Fork the [s3git upstream](https://github.com/s3git/s3git/fork) repository to your own personal repository. Copy the URL and pass it to the ``go get`` command. Go uses git to clone a copy into your project workspace folder.

```sh
$ mkdir -p $GOPATH/src/github.com/s3git
$ cd $GOPATH/src/github.com/s3git
$ git clone https://github.com/$USER_ID/s3git
$ cd s3git
```

Compiling from source
---------------------

Run ``go install`` to build the executable.

```sh
$ go install
```

Setting up git remote as ``upstream``
-------------------------------------

```sh
$ cd $GOPATH/src/github.com/s3git/s3git
$ git remote add upstream https://github.com/s3git/s3git
$ git fetch upstream
$ git merge upstream/master
...
...
$ go install
```

Developer Guidelines
--------------------

``s3git`` community welcomes your contribution. To make the process as seamless as possible, we ask for the following:
* Go ahead and fork the project and make your changes. We encourage pull requests to discuss code changes.
    - Fork it
    - Create your feature branch (git checkout -b my-new-feature)
    - Commit your changes (git commit -am 'Add some feature')
    - Push to the branch (git push origin my-new-feature)
    - Create new Pull Request

* When you're ready to create a pull request, be sure to:
    - Have test cases for the new code. If you have questions about how to do it, please ask in your pull request.
    - Run `go test`
    - Squash your commits into a single commit. `git rebase -i`. It's okay to force update your pull request.
    - Make sure `go test -race ./...` and `go build` completes.

* Read [Effective Go](https://github.com/golang/go/wiki/CodeReviewComments) article from Golang project
    - `s3git` project is fully conformant with Golang style
    - if you happen to observe offending code, please feel free to send a pull request