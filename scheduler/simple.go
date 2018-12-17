package scheduler

import "github.com/zhiyxu/golearn/project/crawler-concurrent/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	//Scheduler Version 1
	//s.workerChan <- request
	//Scheduler Version 2
	//写成goroutine的方式submit，能够快速返回，不会阻塞
	go func() { s.workerChan <- request }()
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(
	c chan engine.Request) {
	s.workerChan = c
}
