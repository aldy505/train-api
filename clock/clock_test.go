package clock_test

import (
	"fmt"
	"testing"
	"time"

	"train-api/clock"
)

func TestNow(t *testing.T) {
	timeNow := time.Now().Round(time.Second)
	now := clock.Now()

	hour, minute, second := timeNow.Clock()
	if now.String() != fmt.Sprintf("%02d:%02d:%02d", hour, minute, second) {
		t.Errorf("expecting clock now to be equal with current time")
	}
}

func TestParse(t *testing.T) {
	t.Run("InvalidSecond", func(t *testing.T) {
		_, err := clock.Parse("00:00:foo")
		if err == nil {
			t.Errorf("expecting an error, got nil")
		}

		if err.Error() != "invalid Second format: foo" {
			t.Errorf("expecting an error of 'invalid Second format: foo', instead got %s", err.Error())
		}
	})

	t.Run("InvalidMinute", func(t *testing.T) {
		_, err := clock.Parse("00:foo")
		if err == nil {
			t.Errorf("expecting an error, got nil")
		}

		if err.Error() != "invalid Minute format: foo" {
			t.Errorf("expecting an error of 'invalid Minute format: foo', instead got %s", err.Error())
		}
	})

	t.Run("InvalidHour", func(t *testing.T) {
		_, err := clock.Parse("foo")
		if err == nil {
			t.Errorf("expecting an error, got nil")
		}

		if err.Error() != "invalid Hour format: foo" {
			t.Errorf("expecting an error of 'invalid Hour format: foo', instead got %s", err.Error())
		}
	})

	t.Run("InvalidSecondRange", func(t *testing.T) {
		_, err := clock.Parse("00:00:90")
		if err == nil {
			t.Errorf("expecting an error, got nil")
		}

		if err.Error() != "Second must be between 0-59" {
			t.Errorf("expecting an error of 'Second must be between 0-59', instead got %s", err.Error())
		}
	})

	t.Run("InvalidMinuteRange", func(t *testing.T) {
		_, err := clock.Parse("00:90")
		if err == nil {
			t.Errorf("expecting an error, got nil")
		}

		if err.Error() != "Minute must be between 0-59" {
			t.Errorf("expecting an error of 'Minute must be between 0-59', instead got %s", err.Error())
		}
	})

	t.Run("InvalidHourRange", func(t *testing.T) {
		_, err := clock.Parse("90")
		if err == nil {
			t.Errorf("expecting an error, got nil")
		}

		if err.Error() != "Hour must be between 0-23" {
			t.Errorf("expecting an error of 'Hour must be between 0-23', instead got %s", err.Error())
		}
	})

	t.Run("EmptyInput", func(t *testing.T) {
		c, err := clock.Parse("")
		if err != nil {
			t.Errorf("unexpected error: %s", err.Error())
		}

		if c.String() != "00:00:00" {
			t.Errorf("expected empty string input to have zero midnight output")
		}
	})

	t.Run("Normal", func(t *testing.T) {
		c, err := clock.Parse("15:04:05")
		if err != nil {
			t.Errorf("unexpected error: %s", err.Error())
		}

		if c.String() != "15:04:05" {
			t.Errorf("expected clock to be '15:04:05', instead got %s", c.String())
		}
	})
}

func TestClock_Before(t *testing.T) {
	t.Run("Equal", func(t *testing.T) {
		c := clock.Now()

		if c.Before(c) {
			t.Error("expecting false")
		}
	})

	t.Run("SecondDifference", func(t *testing.T) {
		c1, _ := clock.Parse("00:00:01")
		c2, _ := clock.Parse("00:00:02")

		if !c1.Before(c2) {
			t.Error("expecting true")
		}

		if c2.Before(c1) {
			t.Error("expecting false")
		}
	})

	t.Run("MinuteDifference", func(t *testing.T) {
		c1, _ := clock.Parse("00:00:01")
		c2, _ := clock.Parse("00:01:02")

		if !c1.Before(c2) {
			t.Error("expecting true")
		}

		if c2.Before(c1) {
			t.Error("expecting false")
		}
	})

	t.Run("HourDifference", func(t *testing.T) {
		c1, _ := clock.Parse("00:00:01")
		c2, _ := clock.Parse("01:01:02")

		if !c1.Before(c2) {
			t.Error("expecting true")
		}

		if c2.Before(c1) {
			t.Error("expecting false")
		}
	})
}

func TestClock_Equal(t *testing.T) {
	c1, _ := clock.Parse("23:59:59")
	c2, _ := clock.Parse("00:00:00")

	if !c1.Equal(c1) {
		t.Error("expecting true")
	}

	if c2.Equal(c1) {
		t.Error("expecting false")
	}
}
