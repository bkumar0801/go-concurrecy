# Concurrency In GO

Go is designed with concurrency in mind and allows us to build complex concurrent pipelines. Go’s approach to concurrency can be best phrased by

<i> Do not communicate by sharing memory; instead, share memory by communicating.</i>

In any other mainstream programming language, when concurrent threads need to share data in order to communicate, a lock is applied to the piece of memory. This typically causes all sorts of issues with race conditions, memory management etc. Instead of applying a lock on the shared variable, Go allows you to communicate (or send) the value stored in that variable from one thread to another. The default behavior is that both the thread sending the data and the one receiving the data will wait till the value reaches its destination. The “waiting” of the threads forces proper synchronization between threads when data is being exchanged.
