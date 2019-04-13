package scheduler

import "imooc.com/Google资深工程师深度讲解Go语言/爬虫/crawier/engine"

// 简单调度器

type SimpleScheduler struct{
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	// send request down to worker chan
	go func() { s.workerChan <- r }()
}

