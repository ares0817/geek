package main

type MyCircularDeque struct {
	cap int
	front int
	last int
	data []int
	count int
}


/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	queue := MyCircularDeque{
		cap: k,
		front: -1,
		last: -1,
		data: make([]int, k),
		count: 0,
	}

	return queue
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.count == 0 {
		this.data[0] = value
		this.front = 0
		this.last = 0
		this.count++
		return true
	}

	if this.IsFull() {
		return false
	}

	this.front = (this.front + this.cap - 1) % this.cap

	this.data[this.front] = value
	this.count++
	return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.count == 0 {
		this.data[0] = value
		this.front = 0
		this.last = 0
		this.count++
		return true
	}

	if this.IsFull() {
		return false
	}

	this.last = (this.last + 1) % this.cap
	this.data[this.last] = value
	this.count++

	return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {

	if this.count == 0 {
		return false
	}

	this.front = (this.front + 1) % this.cap
	this.count--
	return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.count == 0 {
		return false
	}

	this.last = (this.last + this.cap - 1) % this.cap
	this.count--
	return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.count == 0 {
		return -1
	}
	return this.data[this.front]
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.count == 0 {
		return -1
	}
	return this.data[this.last]
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.count == 0
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.count == this.cap
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */