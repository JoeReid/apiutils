package apiutils

import (
	"fmt"
	"net/http"
	"strconv"
)

func intFromReq(r *http.Request, key string) (n int, ok bool, err error) {
	q := r.URL.Query()
	s := q.Get(key)
	if s == "" {
		return 0, false, nil
	}

	// parse the int data
	n, err = strconv.Atoi(s)
	if err != nil {
		return 0, false, err
	}
	return n, true, nil
}

type pagOpts struct {
	count    int
	countSet bool

	skip    int
	skipSet bool
}

func Paginate(r *http.Request, opts ...func(*pagOpts) error) (count int, skip int, err error) {
	const (
		countKey = "count"
		skipKey  = "skip"
	)

	po := pagOpts{}

	po.count, po.countSet, err = intFromReq(r, countKey)
	if err != nil {
		return 0, 0, err
	}

	po.skip, po.skipSet, err = intFromReq(r, skipKey)
	if err != nil {
		return 0, 0, err
	}

	for _, opt := range opts {
		if err = opt(&po); err != nil {
			return 0, 0, err
		}
	}

	if po.count <= 0 {
		return 0, 0, fmt.Errorf("paginate error: count %d must be > 0", po.count)
	}
	if po.skip < 0 {
		return 0, 0, fmt.Errorf("paginate error: skip %d must be >= 0", po.skip)
	}
	return po.count, po.skip, nil
}

func MaxCount(n int) func(*pagOpts) error {
	return func(p *pagOpts) error {
		if p.count > n {
			return fmt.Errorf("paginate error: count %d must be < %d", p.count, n)
		}
		return nil
	}
}

func DefaultCount(n int) func(*pagOpts) error {
	return func(p *pagOpts) error {
		if !p.countSet {
			p.count = n
		}
		return nil
	}
}

func MaxSkip(n int) func(*pagOpts) error {
	return func(p *pagOpts) error {
		if p.skip > n {
			return fmt.Errorf("paginate error: skip %d must be < %d", p.skip, n)
		}
		return nil
	}
}

func DefaultSkip(n int) func(*pagOpts) error {
	return func(p *pagOpts) error {
		if !p.skipSet {
			p.skip = n
		}
		return nil
	}
}
