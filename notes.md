# Parallel Computing and Concurrency in Go

## Amdahl's Law
Restriction of speed up a function of ratio of parallel to sequential hardware architecture.

## Gustafson's Law
If we can increase problem size, we can get linear speed up wrt to processor units. If we have idling processor units, then perhaps we can use them to do something else!

## Single Core Concurrency
A single core can mimick benefits of concurrency via scheduling. For example, a task that is awaiting I/O can be kept on wait queue and other task on the execution queue is given processor time. As soon as the I/O is completed, the waiting task is moved to the execution queue.

## Processes, Threads and Green Threads


