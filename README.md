# Concurrency In GO

Go is designed with concurrency in mind and allows us to build complex concurrent pipelines. Go’s approach to concurrency can be best phrased by

<i> Do not communicate by sharing memory; instead, share memory by communicating.</i>

In any other mainstream programming language, when concurrent threads need to share data in order to communicate, a lock is applied to the piece of memory (this is also an overhead on programmer to select right locking mechanism). This typically causes issues like race conditions / deadlock, memory management etc. Instead of applying a lock on the shared memory (critical section), Go allows you to communicate (or send) the value from one routine to another. The default behavior is that both the routines sending the data and the one receiving the data will wait till the value reaches its destination. The “waiting” of the routines forces proper synchronization between routines when data is being exchanged.
