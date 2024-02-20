package goprogress

import (
	"sync"
)

type Progress struct {
	totalUnits     int64
	completedUnits int64

	mutex *sync.Mutex
}

func (p *Progress) IsCompleted() bool {
	return p.totalUnits == p.completedUnits
}

func (p *Progress) GetCompletedUnits() int64 {
	return p.completedUnits
}

func (p *Progress) IncrementProgress(units int64) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if (p.completedUnits + units) > p.totalUnits {
		return ErrorProgressOverTotalCount
	}

	p.completedUnits += units

	return nil
}

func (p *Progress) FractionCompleted() float64 {
	if p.IsCompleted() {
		return 1.0
	}

	return float64(p.completedUnits) / float64(p.totalUnits)
}

func NewProgress() Progress {
	return Progress{mutex: &sync.Mutex{}}
}

func NewProgressWithUnits(units int64) Progress {
	return Progress{
		totalUnits:     units,
		completedUnits: 0,
		mutex:          &sync.Mutex{},
	}
}
