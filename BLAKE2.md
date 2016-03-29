BLAKE2
======

The [BLAKE2](https://blake2.net/blake2.pdf) hashing algorithm is a fast and modern hashing algorithm that is optimized for speed in software. BLAKE2 has a comprehensive tree-hashing mode that lends itself very well for parallel implementations and has some interesting properties that are explained below. On top of that it is also nicely suitable for GPU implementations allowing for even greater performance improvements.

BLAKE2 Tree mode/Unlimited fanout
---------------------------------

In addition to the 'normal' sequential mode that most hashing algorithms use, BLAKE2 has a very flexible tree-hashing mode. Although BLAKE2 supports arbitrary-depth trees, s3git uses a special mode called **unlimited fanout** as shown here:

```
                  /=====\
                  | 1:0 |
                  \=====/

/-----\  /-----\  /-----\  /-----\      /=====\
| 0:0 |  | 0:1 |  | 0:2 |  | 0:3 |  ... | 0:N | 
\-----/  \-----/  \-----/  \-----/      \=====/
```

In this diagram the boxes represent leaves whereby the label `i:j` represents a node's depth `i` and offset `j`. Double-lined nodes (including leaves) are the last nodes of a layer. The leaves process chunks of data of `leaf length` bytes independently of each other, and subsequently the root node hashes the concatenation of the hashes of the leaves.

For BLAKE2's unlimited fanout mode the depth is always fixed at 2 and there can be as many leaves as are required given the size of the input. Note that the `node offset` and `node depth` parameters ensure that each invocation of BLAKE2 uses a different hash function (and hence will generate a different output for the same input).
 
BLAKE2 in s3git 
---------------

Let's see an example of how s3git uses BLAKE2

```
$ s3git init
Initialized empty s3git repository
$ echo 'hello s3git' | s3git add
Added: 18e622875a89cede0d7019b2c8afecf8928c21eac18ec51e38a8e6b829b82c3ef306dec34227929fa77b1c7c329b3d4e50ed9e72dc4dc885be0932d3f28d7053
$ more .s3git/stage/46/dd/d7b91748c4d253e328a9644d78b3e3a298ebbbab462891502f05e956ef7ec03c8e0978e5160a858cc50ca6b37176248b602d50d0c609abe75b462b6dddcc
hello s3git
```

(As you may have figured, s3git stores content that is not yet committed and pushed in the `.s3git/stage` directory, see also `s3git status`)

For the input `hello s3git` the (sole) leaf hash is `46ddd7b91748c4d253e328a9644d78b3e3a298ebbbab462891502f05e956ef7ec03c8e0978e5160a858cc50ca6b37176248b602d50d0c609abe75b462b6dddcc` which (when processed as bytes) results in `18e622875a89cede0d7019b2c8afecf8928c21eac18ec51e38a8e6b829b82c3ef306dec34227929fa77b1c7c329b3d4e50ed9e72dc4dc885be0932d3f28d7053` as the root hash.

Let's try a larger file of 8 MB that will give more than one leaf (s3git's default leaf size is 5 MB)

```
$ dd if=/dev/zero bs=1048576 count=8 | s3git add
Added: 2039f91853e3cf31ae3d587609d0459331b35863a743cb3ef9c4e2baf26bb317e2e7f06b594285c97e58c47750b29efebca93e63dd24e1424737e6664ade7414
$ ls -l .s3git/stage/30/21/a7f3d7ed2ac353fa380ebfacb3e8e2e8e4ebfb1b28d24a56d3bd79d715470edc3ca868576a4d17dae886b61ba72bcd3780b67a3d1be1c9cb1b25d7cd1a61
-rw-r--r--  1 frankw  staff  5242880 Mar 14 13:19 .s3git/stage/30/21/a7f3d7ed2ac353fa380ebfacb3e8e2e8e4ebfb1b28d24a56d3bd79d715470edc3ca868576a4d17dae886b61ba72bcd3780b67a3d1be1c9cb1b25d7cd1a61
$ ls -l .s3git/stage/6c/ac/33b4fa6803ae784db76e4a8b43c074a7fcdf2dc4cce558cc01c5ff6f909a6fb3fa5e56b7205aa4b4c74a70545c20fce09f2b85edefbc43e39507f21ea356
-rw-r--r--  1 frankw  staff  3145728 Mar 14 13:19 .s3git/stage/6c/ac/33b4fa6803ae784db76e4a8b43c074a7fcdf2dc4cce558cc01c5ff6f909a6fb3fa5e56b7205aa4b4c74a70545c20fce09f2b85edefbc43e39507f21ea356
```

As you can see two new leaves have been created with sizes of 5242880 and 3145728 respectively. If you concatenate `3021a7f3d7ed2ac353fa380ebfacb3e8e2e8e4ebfb1b28d24a56d3bd79d715470edc3ca868576a4d17dae886b61ba72bcd3780b67a3d1be1c9cb1b25d7cd1a61` and `6cac33b4fa6803ae784db76e4a8b43c074a7fcdf2dc4cce558cc01c5ff6f909a6fb3fa5e56b7205aa4b4c74a70545c20fce09f2b85edefbc43e39507f21ea356` as bytes you will exactly get `2039f91853e3cf31ae3d587609d0459331b35863a743cb3ef9c4e2baf26bb317e2e7f06b594285c97e58c47750b29efebca93e63dd24e1424737e6664ade7414` as the root hash.

Now this is all nice and you may be wondering why this matters but let's dive into the next section for this.

Cloud storage
-------------

s3git allows you to commit your changes and push them into cloud storage like S3. There are two formats how to do this:

-  **deduplicated** (also called deduped)
-  **hydrated** (or concatenated)

Let's examine both in the following examples.

### Deduplicated

For the deduplicated format multiple files or objects are created:

-  the leaf nodes are stored under their own hash
-  the list of hashes of the leaves is stored under the name of the root hash

For the two files that were added above this generates the following result:

```
$ s3git push
Done.
$ aws s3 ls s3://s3git-test
2016-03-21 20:05:21         64 18e622875a89cede0d7019b2c8afecf8928c21eac18ec51e38a8e6b829b82c3ef306dec34227929fa77b1c7c329b3d4e50ed9e72dc4dc885be0932d3f28d7053
2016-03-21 20:05:21        128 2039f91853e3cf31ae3d587609d0459331b35863a743cb3ef9c4e2baf26bb317e2e7f06b594285c97e58c47750b29efebca93e63dd24e1424737e6664ade7414
2016-03-21 20:05:21    5242880 3021a7f3d7ed2ac353fa380ebfacb3e8e2e8e4ebfb1b28d24a56d3bd79d715470edc3ca868576a4d17dae886b61ba72bcd3780b67a3d1be1c9cb1b25d7cd1a61
2016-03-21 20:05:21         12 46ddd7b91748c4d253e328a9644d78b3e3a298ebbbab462891502f05e956ef7ec03c8e0978e5160a858cc50ca6b37176248b602d50d0c609abe75b462b6dddcc
2016-03-21 20:05:21    3145728 6cac33b4fa6803ae784db76e4a8b43c074a7fcdf2dc4cce558cc01c5ff6f909a6fb3fa5e56b7205aa4b4c74a70545c20fce09f2b85edefbc43e39507f21ea356
```

As you can see the first two objects are the root hashes for the two streams. The first object of size 64 bytes contains a single pointer whereas the second object of size 128 bytes holds 2 pointers. The remaining three objects are all leaf nodes and contain the contents.

### Hydrated

For the hydrated format just a single file or object is created that stores the full original contents. (This naturally prevents deduplication and indeed hydrated can be thought of as un-deduplicated.)

For the two files that were added above this generates the following result:

```
$ s3git push --hydrate
Done.
$ aws s3 ls s3://s3git-test
2016-03-21 20:06:43         12 18e622875a89cede0d7019b2c8afecf8928c21eac18ec51e38a8e6b829b82c3ef306dec34227929fa77b1c7c329b3d4e50ed9e72dc4dc885be0932d3f28d7053
2016-03-21 20:06:43    8388608 2039f91853e3cf31ae3d587609d0459331b35863a743cb3ef9c4e2baf26bb317e2e7f06b594285c97e58c47750b29efebca93e63dd24e1424737e6664ade7414
```

It will be obvious that the first object stores the first example whereas the second object stores the second example.

### Deduped vs hydrated

Both methods have their advantages and disadvantages as is listed in this table:

|               | Deduped | Hydrated |
| ------------- |:-------:| :-------:|
| Deduplication |    ✓    |          |
| Direct access |         |     ✓    |
| Rolling hash  |    ✓    |          |
| Encryption    |    ✓    |     ✓    |

Deduplication obviously saves on storage and bandwidth costs as duplicate content is stored just once. s3git deduplicates at the repository level so it is not just 'file-level' or 'block-level' deduplication but global data deduplication within a repository. When combined with rolling hashes it allows for even better data reduction levels and a nice level of resistance to content changes anywhere within the input streams. A future blog will elaborate on this.

For certain content (like huge videos or input files for mapreduce jobs) the direct access feature of the hydrated format is a nice benefit. It allows content to be accessed and fetched directly out of cloud storage using a single (presigned) URL without the need for any intermediate server step that glues the leaf chunks together.

Also note that for encrypted content you might just as well use the hydrated format since the very nature of encryption avoids duplicate content (or your encryption is broken...).

### Wait, bad idea! Store different content under same name?!

Now you may wonder that it is a bad idea to store different content under the same name as this might create confusion. When using a normal hashing algorithm (without tree-hashing support) that is the case, but for BLAKE2 this is not a problem.

In order to determine whether a root hash stores a deduped object, the following must be true:
- The size needs to be a multiple of 64 bytes (which allows for a quick check), and
- When hashed at level 1 and as last node it must return its own root hash.

BLAKE2 even protects you (via the level property `i`) from content that (by chance) exactly matches the content that is used to compute any root hash. As an example, imagine an input file that has content that fully matches the (sole) leaf hash for the `hello s3git` case above. 

```
$ echo 46ddd7b91748c4d253e328a9644d78b3e3a298ebbbab462891502f05e956ef7ec03c8e0978e5160a858cc50ca6b37176248b602d50d0c609abe75b462b6dddcc | xxd -r -p | s3git add
Added: 4cba3e9d94f5c2a643ee365487249342e16d8e58cfd53c7b2022b7472b46cd30b08af32db1998a9f93a029bd086e4b1b744af2b46c54fab106beadb3b4cbed78
$ hexdump -C .s3git/stage/c7/88/1bd31c1d13ac080ce7188d92fc7296411e27df641c0431c305b299108b8c2c09c68076a760feee685a66b9cf70b45954f24191bc02497a1de338c76d91a8
00000000  46 dd d7 b9 17 48 c4 d2  53 e3 28 a9 64 4d 78 b3  |F....H..S.(.dMx.|
00000010  e3 a2 98 eb bb ab 46 28  91 50 2f 05 e9 56 ef 7e  |......F(.P/..V.~|
00000020  c0 3c 8e 09 78 e5 16 0a  85 8c c5 0c a6 b3 71 76  |.<..x.........qv|
00000030  24 8b 60 2d 50 d0 c6 09  ab e7 5b 46 2b 6d dd cc  |$.`-P.....[F+m..|
00000040
```

As you can see when this (byte) stream is added to s3git, it generates a leaf hash of `c7881bd31c1d13ac080ce7188d92fc7296411e27df641c0431c305b299108b8c2c09c68076a760feee685a66b9cf70b45954f24191bc02497a1de338c76d91a8` which in turn leads to a root hash of `4cba3e9d94f5c2a643ee365487249342e16d8e58cfd53c7b2022b7472b46cd30b08af32db1998a9f93a029bd086e4b1b744af2b46c54fab106beadb3b4cbed78`.

And luckily this is not quite the same as the root hash `18e622875a89cede0d7019b2c8afecf8928c21eac18ec51e38a8e6b829b82c3ef306dec34227929fa77b1c7c329b3d4e50ed9e72dc4dc885be0932d3f28d7053` that you will get for `hello s3git`.

### Mix deduped and hydrated in same repo

Due to the fact that we can determine whether content is stored deduped or not, it is even possible to mix both deduped and hydrated storage in the same repository. Or in the case of multiple remotes (as will be supported by s3git) you can store it deduped in one remote and hydrated in another.

For cloud drives such as Amazon Cloud Drive this is a nice capability since by storing for example videos in hydrated format, things like thumbnails and viewing are possible. 

Conclusion
----------

Using BLAKE2 in the unlimited fanout mode allows s3git to seamlessy support both deduplication as well as hydration within the same repository.

Due to the hashing nature it is virtually certain (at least within our lifetimes (and those of our grand children)) that no hash collisions will occur. More on this in a future blog.

References
----------

- [BLAKE2: simpler, smaller, fast as MD5](https://blake2.net/blake2.pdf)
