package queue

import "errors"

func (queue *Queue) Len() int {
	return len(*queue)
}

func (queue *Queue) IsEmpty() bool {
	return len(*queue) == 0
}

func (queue *Queue) Cap() int {
	return cap(*queue)
}

func (queue *Queue) Push(value interface{}) {
	*queue = append(*queue, value)
}

func (queue *Queue) Pop() (interface{}, error) {
	theQueue := *queue
	if len(theQueue) == 0 {
		return nil, errors.New("out of the index, len is 0")
	}
	value := theQueue[0]
	*queue = theQueue[1 : len(theQueue)-1]
	return value, nil
}
