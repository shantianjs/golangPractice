package practice

type MyCircularQueue struct {
	data []int
	head, length, capacity int
}


/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		make([]int,k,k),
		0, 0,k,
	}
}


/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.data[(this.head + this.length) % this.capacity] = value
	this.length++

	return true
}


/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty(){
		return false
	}
	if this.head == this.capacity - 1 {
		this.head = 0
	} else {
		this.head += 1
	}
	this.length--
	return true
}


/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.length == 0 {
		return -1
	}
	return this.data[this.head]
}


/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.length == 0 {
		return -1
	}
	return this.data[(this.head + this.length - 1) % this.capacity ]
}


/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.length == 0
}


/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.length == this.capacity
}

