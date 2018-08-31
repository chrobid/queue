package queue

// Node is an element in the queue
type node struct {
	data int
	next *node
	prev *node
}

// Queue implements a queue data structure. This queue can only hold integers.
type Queue struct {
	head *node
	tail *node
}

// NewQueue instantiates an empty queue
func NewQueue() *Queue {
	return &Queue{}
}

// Push adds a new node to the tail
func (q *Queue) Push(i int) {
	newNode := &node{}
	newNode.data = i
	if q.tail == nil {
		q.head, q.tail = newNode, newNode
	} else {
		q.tail.prev = newNode
		newNode.next = q.tail
		q.tail = newNode
	}
}

// Pop returns the head from the queue, and removes it from the queue
func (q *Queue) Pop() (i int) {
	// if you pop from an empty queue, it returns zero. I'm not sure if this behavior is OK.
	if q.head == nil {
		return
	} else if q.head == q.tail {
		i = q.head.data
		q.head, q.tail = nil, nil
	} else {
		i = q.head.data
		q.head = q.head.prev
		q.head.next = nil
	}
	return
}

// IsEmpty returns true if the queue is empty, and false if it's not
func (q *Queue) IsEmpty() bool {
	if q.head == nil {
		return true
	}
	return false
}

// ToArray returns an array containing all nodes in the queue
func (q *Queue) ToArray() []int {
	var nodes []int
	currentNode := q.tail
	for currentNode != nil {
		nodes = append(nodes, currentNode.data)
		currentNode = currentNode.next
	}
	return nodes
}
