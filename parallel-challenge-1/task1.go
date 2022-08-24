package training

import (
	"context"
)

type result struct {
	ok  bool
	err error
}

func RunParallel(ctx context.Context, tasks ...func(ctx context.Context) (bool, error)) (bool, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	ch := make(chan result)

	for _, task := range tasks {
		go func(task func(ctx context.Context) (bool, error)) {
			ok, err := task(ctx)
			if err != nil {
				cancel()
			}
			ch <- result{
				ok:  ok,
				err: err,
			}
		}(task)
	}

	// Since we are looping over the tasks, we don't need a WaitGroup because this loop will end
	// as long as all the tasks return.
	for range tasks {
		res := <-ch
		if res.err != nil {
			return false, res.err
		}
		if !res.ok {
			return false, nil
		}
	}

	return true, nil
}
