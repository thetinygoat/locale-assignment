package queue

import (
	"github.com/adjust/redismq"
)

// NewQueue instantiates a new queue
func NewQueue(name string) *redismq.Queue {

	return redismq.CreateQueue("localhost", "6379", "", 9, name)

}
