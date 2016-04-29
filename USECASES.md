
Use cases for s3git
===================

Below is a list of various use cases where s3git could prove useful.

Build and Release Management
----------------------------

- For continuous deployment store the collection of all your binaries, source files and config
- Commit this as a single release and push into cloud
- Allow servers to checkout a release and go live
- Machines can selectively fall back to any previous version if needed

See [binary release management](https://github.com/s3git/s3git/blob/master/BINARY-RELEASE-MANAGEMENT.md) for an example using the Kubernetes project.

DevOps Scenarios
----------------

- Have a single integrated storage platform/repository for production data and other uses
- Fast clone of versioned copy of production data for specific development or testing
- Work with a subset of huge repositories, even on small local SSDs 
- Fully programmable for automatation of tasks

Data Consolidation
------------------

- Combine different datasets or data silos into a single repository
- Infinitely scalable
- Leverage deduplication at global scale
- Provide unlimited backup and recovery of data
- Go back in time to previous versions
- Allow replication of data across multiple data centers
- Prevent forklift upgrades

Analytics
---------

- Have many clients and/or apps write, commit and push concurrently to the same repository in the cloud
- Scalable to 100 of millions of files
- Allow analytics using eg MapReduce to access data directly out of cloud storage

Photo and Video storage
-----------------------

- Store and organize all your photos and videos in original quality in the cloud.
- Sync to cloud storage, for example Amazon Cloud Drive to get thumbnails etc.
- Allow to get all content back out in original quality (do a full clone)
- Have integrity of the data (know you have everything and that it has not been messed with)
- Allow incremental pulls (fetch all the new material)
- (this is what the author uses for his own family archive)

Request for other use cases
---------------------------

If you see another use case for s3git, please file an issue and describe it!
