package goprogress_test

import (
	"errors"
	"testing"

	gp "github.com/KiritoCyanPine/go-progress"
)

func TestEmpty(t *testing.T) {
	p := gp.NewProgress()

	if p.IsCompleted() != true {
		t.Error("Progress is empty, but shown as not completed")
	}
}

func TestNotEmpty(t *testing.T) {
	p := gp.NewProgressWithUnits(100)

	if p.IsCompleted() != false {
		t.Error("Progress not used but is completed")
	}
}

func TestProgressIncrementBy1(t *testing.T) {
	p := gp.NewProgressWithUnits(100)

	if err := p.IncrementProgress(1); err != nil {
		t.Error("Progress could not be incremented, due to err ", err)
	}

	if p.GetCompletedUnits() != 1 {
		t.Error("Progress was incremented by 1, but is not 1 now")
	}
}

func TestProgressIncrementBy99(t *testing.T) {
	p := gp.NewProgressWithUnits(100)

	if err := p.IncrementProgress(99); err != nil {
		t.Error("Progress could not be incremented, due to err ", err)
	}

	if p.GetCompletedUnits() != 99 {
		t.Error("Progress was incremented by 99, but is not 99 now")
	}
}

func TestProgressIncrementToTotalCount(t *testing.T) {
	p := gp.NewProgressWithUnits(100)

	if err := p.IncrementProgress(100); err != nil {
		t.Error("Progress could not be incremented, due to err ", err)
	}

	if p.GetCompletedUnits() != 100 {
		t.Error("Progress was incremented by 100, but is not 100 now")
	}

	if p.IsCompleted() != true {
		t.Error("Progress is completed, but shown as not completed")
	}
}

func TestProgressIncrementToTotalCountInSteps(t *testing.T) {
	p := gp.NewProgressWithUnits(100)

	if err := p.IncrementProgress(50); err != nil {
		t.Error("Progress could not be incremented, due to err ", err)
	}

	if p.GetCompletedUnits() != 50 {
		t.Error("Progress was incremented by 50, but is not 50 now")
	}

	if p.IsCompleted() == true {
		t.Error("Progress is not completed, but shown as completed")
	}

	if err := p.IncrementProgress(50); err != nil {
		t.Error("Progress could not be incremented, due to err ", err)
	}

	if p.GetCompletedUnits() != 100 {
		t.Error("Progress was incremented by 100, but is not 100 now")
	}

	if p.IsCompleted() != true {
		t.Error("Progress is comepleted, but shown as not completed")
	}
}

func TestProgressIncrementOverTotalCount(t *testing.T) {
	p := gp.NewProgressWithUnits(100)

	if err := p.IncrementProgress(101); err == nil {
		t.Error("Progress should not be able to excede totalCount ")
	} else {
		if !errors.Is(err, gp.ErrorProgressOverTotalCount) {
			t.Error("Error should be ", gp.ErrorProgressOverTotalCount.Error())
		}
	}
}

func TestEmptyProgressFractionWorkCompleted(t *testing.T) {
	p := gp.NewProgress()

	if p.FractionCompleted() != 1.0 {
		t.Error("empty progress fraction completed should be 1")
	}
}

func TestProgressWithCountFractionWorkCompletedAndNoProgress(t *testing.T) {
	p := gp.NewProgressWithUnits(100)

	if p.FractionCompleted() != 0.0 {
		t.Error("progress fraction completed should be 0.0")
	}
}

func TestProgressWithCount50PercentFractionCompleted(t *testing.T) {
	p := gp.NewProgressWithUnits(100)

	if err := p.IncrementProgress(50); err != nil {
		t.Error("Progress could not be incremented, due to err ", err)
	}

	if p.FractionCompleted() != 0.5 {
		t.Error("progress fraction completed should be 0.5 for 50 percent work done")
	}
}

func TestProgressFractionCompletedratio(t *testing.T) {

	var totalUnits int64 = 100

	p := gp.NewProgressWithUnits(totalUnits)

	for i := int64(1); i <= totalUnits; i++ {
		if err := p.IncrementProgress(1); err != nil {
			t.Error("Progress could not be incremented, due to err ", err)
			t.FailNow()
		}

		if p.FractionCompleted() != (float64(i) / float64(totalUnits)) {
			t.Errorf("Expected fraction completed as %f , but got %f", float64(i/totalUnits), p.FractionCompleted())
			t.FailNow()
		}
	}
}
