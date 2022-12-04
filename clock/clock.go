package clock

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Clock struct {
	Hour   int
	Minute int
	Second int
}

func Now() Clock {
	now := time.Now().Round(time.Second)

	return Clock{
		Hour:   now.Hour(),
		Minute: now.Minute(),
		Second: now.Second(),
	}
}

func Parse(s string) (Clock, error) {
	var c Clock
	if s == "" {
		return c, nil
	}

	parts := strings.SplitN(s, ":", 3)
	if len(parts) >= 3 {
		second, err := strconv.Atoi(parts[2])
		if err != nil {
			return Clock{}, fmt.Errorf("invalid Second format: %s", parts[2])
		}

		if second < 0 || second > 59 {
			return Clock{}, fmt.Errorf("Second must be between 0-59")
		}

		c.Second = second
	}

	if len(parts) >= 2 {
		minute, err := strconv.Atoi(parts[1])
		if err != nil {
			return Clock{}, fmt.Errorf("invalid Minute format: %s", parts[1])
		}

		if minute < 0 || minute > 59 {
			return Clock{}, fmt.Errorf("Minute must be between 0-59")
		}

		c.Minute = minute
	}

	if len(parts) >= 1 {
		hour, err := strconv.Atoi(parts[0])
		if err != nil {
			return Clock{}, fmt.Errorf("invalid Hour format: %s", parts[0])
		}

		if hour < 0 || hour > 23 {
			return Clock{}, fmt.Errorf("Hour must be between 0-23")
		}

		c.Hour = hour
	}

	return c, nil
}

func (h Clock) String() string {
	return fmt.Sprintf("%02d:%02d:%02d", h.Hour, h.Minute, h.Second)
}

func (h Clock) Before(c Clock) bool {
	if h.Hour == c.Hour {
		if h.Minute == c.Minute {
			if h.Second == c.Second {
				return false
			}

			return h.Second < c.Second
		}

		return h.Minute < c.Minute
	}

	return h.Hour < c.Hour
}

func (h Clock) After(c Clock) bool {
	if h.Hour == c.Hour {
		if h.Minute == c.Minute {
			if h.Second == c.Second {
				return false
			}

			return h.Second > c.Second
		}

		return h.Minute > c.Minute
	}

	return h.Hour > c.Hour
}

func (h Clock) Equal(c Clock) bool {
	return h.Hour == c.Hour && h.Minute == c.Minute && h.Second == c.Second
}
