package training

import (
	"context"
	"sync"
)

type result struct {
	ok  bool
	err error
}

func RunParallel(ctx context.Context, tasks ...func() (bool, error)) (bool, error) {
	// Tip: WaitGroup shouldn't be copied. It should only be passed as pointers.
	// We could be passing it to the goroutines below as a pointer but they are available anyway.
	wg := sync.WaitGroup{}
	wg.Add(len(tasks))

	ch := make(chan result)

	for _, task := range tasks {
		go func(task func() (bool, error)) {
			ok, err := task()
			ch <- result{
				ok:  ok,
				err: err,
			}
			wg.Done()
		}(task)
	}

	go func() {
		// Wait until all of the tasks are done and then close the channel.
		// Do this in a go function so that we don't block the rest of the function.
		wg.Wait()
		close(ch)
	}()

	for res := range ch {
		if res.err != nil {
			return false, res.err
		}
		if !res.ok {
			return false, nil
		}
	}

	return true, nil
}
