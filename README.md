[![GoDoc](https://godoc.org/github.com/michaelmaltese/T-DFS?status.png)](https://godoc.org/github.com/michaelmaltese/T-DFS) [![Build Status](https://travis-ci.org/michaelmaltese/T-DFS.svg?branch=master)](https://travis-ci.org/michaelmaltese/T-DFS)

Writing a HDFS clone in [Go](http://golang.org) to learn more about Go and the nitty-gritty of distributed systems. 

## Features/TODO

- [x] MetaDataNode/DataNode handle uploads
- [x] MetaDataNode/DataNode handle downloads
- [x] DataNode dynamically registers with MetaDataNode
- [x] DataNode tells MetaDataNode its blocks on startup
- [x] MetaDataNode persists file->blocklist map
- [x] DataNode pipelines uploads to other DataNodes
- [x] MetaDataNode can restart and DataNode will re-register (heartbeats)
- [x] Tell DataNodes to re-register if MetaDataNode doesn't recognize them
- [x] Drop DataNodes when they go down (heartbeats)
- [x] DataNode sends size of data directory (heartbeat)
- [x] MetaDataNode obeys replication factor
- [x] MetaDataNode balances based on current reported space
- [x] MetaDataNode balances based on expected new blocks
- [x] MetaDataNode handles not enough DataNodes for replication
- [x] Have MetaDataNode manage the block size stuff (in HDFS, clients can change this per-file)
- [x] Re-replicate blocks when a DataNode disappears
- [x] Drop over-replicated blocks when a DataNode comes up
- [x] Looking at DataNode utilization should take into account the DeletionIntents and ReplicationIntents
- [x] Grace period for replicating new blocks
- [x] MetaDataNode balances blocks as it runs!
- [x] Record hash digest of blocks, reject send if hash is wrong
- [x] DataNode needs to keep track of blocks it's receiving / deleting / checking so that the integrity checker can run only on real blocks
- [x] Remove blocks if checksum doesn't match
- [x] Run a cluster in a single process for testing
- [x] Structure things better
- [x] Resiliency to weird protocol stuff (run the RPC loop manually?)
- [x] Command line parser doesn't work that well (try "main datanode -help")
- [ ] Events from servers for testing
- [ ] Better configuration handling (defaults)
- [ ] Allow decommissioning nodes
- [ ] Better logging, so warnings normally can be fatal for tests (two levels: warn that this process broke, and warn that somebody we're communicating with broke)
- [ ] Don't need to wait around to delete blocks, just prevent any new reads and we'll come back to them
- [ ] DataNode should do stuff on startup, and then spawn workers, not just spawn everybody (race conditions with address and data directories)
- [ ] Support multiple MetaDataNodes somehow (DHT? Raft? Get rid of MetaDataNodes and use Gossip?)
- [ ] Keep track of MoveIntents (subtract from predicted utilization of node), might fix the volatility when re-balancing
- [ ] HashiCorp claims heartbeats are inefficient (linear work aafo number of nodes). Use Gossip?
- [ ] Don't force a long-running connection for creating a file, give the client a lease and let them re-connect
- [ ] If a client tries to upload a block and every DataNode in its list is down, it needs to get more from the MetaDataNode.
- [ ] Keep track of blocks as we're creating a file, if the client bails before committing then delete the blocks.
