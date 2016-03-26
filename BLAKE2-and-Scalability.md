
BLAKE2 and Scalability
======================

The s3git-100m repo (based on the [Yahoo Flickr Creative Commons 100M](http://aws.amazon.com/public-data-sets/multimedia-commons/) dataset, see [here](https://github.com/s3git/s3git#clone-the-yfcc100m-dataset) how to clone) has served as a nice test bed to do some statistical analysis. When the hashes are sorted it is easy to compute the distance between two consecutive hashes and determine how many leading zero bits the distance has. This then gives us a good measure for how close two neighbouring hashes are to a potential collision.

The results of this analysis are plotted in the following chart:

![BLAKE2 Distance Distribution for YFCC100M dataset](https://s3.amazonaws.com/s3git-assets/BLAKE2-distance-distribution-for-yfcc100m-dataset.png)

As you can see (on the logarithmic scale) it is a pretty linear outcome as you would have liked to get. And almost like 'clockwork' it declines by a factor of two per extra bit which also makes sense intuitively (eg. `63` for `46` bits, `31` for `47` bits, `15` for `48` bits, etc.). For the full details check out the table at the end.

One of the "closest" collisions (still way off...) was found with 52 equal leading zero bits, the equivalent of 13 hexadecimal chars: 

```
$ s3git ls 7d5542e8c4e7d
7d5542e8c4e7d09de7bc3032ad594505fd99e1f62c308f6f64098148893b102fcb0af1d2ee4c818a3611bbc9278f48f97d4b2105a7cc6a2f4f72f07eb60cf8b8
7d5542e8c4e7dd5e097024d8be86339ef9060200cac46c08a83b9a3a31a3c37b294fc1c3083db6996f2883e3386f8c880195b1bd6178390a22647882549747e8
```

Collision Course?
-----------------

Given the maximum of 52 equal leading bits that we found in the 100M dataset, that still leaves 460 bits that are different out of the 512 bits total.

If we assume that with a doubling of the size of the dataset we would get an extra bit of equality at the front then we can create the following table:

| Dataset size | Leading zero bits |
|:-----------:| :-------:|
| 100M | 52 |
| 200M | 53 |
| 400M | 54 |
| 800M | 55 |
| ... | ... |
| 2.6E69 | 256 |
| ... | ... |
| 3.0E146 | 512 |

So there we have it: at an expected `3.0E146` objects there will be a collision. However according to [wikipedia](https://en.wikipedia.org/wiki/Observable_universe#Matter_content_.E2.80.93_number_of_atoms) there are about `10E80` atoms in the universe so that is going to be a bit of a challenge.

Conclusion
----------

As you can see BLAKE2 is doing a pretty nice job of distributing the hashes within its space for the YFCC100M dataset.

So we can be pretty confident that we will not likely see a collision any time soon. Sleep well tonight.

Table
-----

This is the table used to plot the chart above.

| Leading zero bits | Appearance |
|:-----------:| -------:|
| 22 | 828 |
| 23 | 284418 |
| 24 | 4999843 |
| 25 | 17466963 |
| 26 | 24464488 |
| 27 | 20797985 |
| 28 | 13610473 |
| 29 | 7804585 |
| 30 | 4176794 |
| 31 | 2159028 |
| 32 | 1097765 |
| 33 | 553236 |
| 34 | 279060 |
| 35 | 139778 |
| 36 | 69909 |
| 37 | 34682 |
| 38 | 17505 |
| 39 | 8668 |
| 40 | 4366 |
| 41 | 2190 |
| 42 | 1127 |
| 43 | 526 |
| 44 | 255 |
| 45 | 150 |
| 46 | 63 |
| 47 | 31 |
| 48 | 15 |
| 49 | 8 |
| 50 | 6 |
| 51 | 2 |
| 52 | 2 |