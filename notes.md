# Parallel Computing and Concurrency in Go

## Amdahl's Law
Restriction of speed up a function of ratio of parallel to sequential hardware architecture.

## Gustafson's Law
If we can increase problem size, we can get linear speed up wrt to processor units. If we have idling processor units, then perhaps we can use them to do something else!

## Single Core Concurrency
A single core can mimick benefits of concurrency via scheduling. For example, a task that is awaiting I/O can be kept on wait queue and other task on the execution queue is given processor time. As soon as the I/O is completed, the waiting task is moved to the execution queue.

## Processes, Threads and Green Threads

Each process has its own memory space. One of them crashing does not affect the other. They are costly to create.

Threads share memory space, so it is faster to create threads. Threads are scheduled by OS kernel (kernel level threads). Kernel spends some time in choosing and queuing threads - context switching. Normally it is small, however we do not want it to dominate the total execution time.

Green threads are user level threads, where the context switching or scheduling is managed at the program level, so the kernel faces an "optimal" number of threads.

One particular scenario to be aware of: suppose each os thread has a number of green threads and if one of the green thread in waiting for a sync op to finish, the entire os thread will be moved to wait queue, even though other green threads can execute.

Go uses a hybrid pattern to mitigate this issue by moving those threads into an os thread so they can continue executing.

## Interprocess Communication
1. Shared Memory
    - ease of implementation
    - multiple variables can be shared
2. Message Passing

