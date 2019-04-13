package engine

import "log"

type ConcurrentEngine struct {
	Scheduler Scheduler // 调度器
	WorkCount int 		// 工作数
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
} 

func (e *ConcurrentEngine) Run (seeds ...Request)  {
	// 初始化输入输出
	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkerChan(in)

	for i := 0; i < e.WorkCount; i++ {
		createWorker(in, out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	// 接收打印
	itomCount := 0
	for {
		result := <- out
		for _, item := range result.Items {
			log.Printf("Got item %d； %v", itomCount, item)
			itomCount++
		}

		for _, request := range result.Request {
			e.Scheduler.Submit(request)
		}
	}
}
func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for  {
			request := <- in
			result, err := work(request)
			if err != nil{
				continue
			}

			out <- result
		}

	}()
}