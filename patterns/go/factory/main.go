package main

import (
	"sync"
	"time"
)

type LogMode string

const (
	LowMode       LogMode = "low"
	MediumMode    LogMode = "medium"
	HighMode      LogMode = "high"
	AlwaysMode    LogMode = "always"
	ForbiddenMode LogMode = "forbidden"
)

// In TriggerLogDuration, if error times < TriggerLogCount pass, else print error log.
type LogCounter struct {
	FirstLogTime       time.Time
	LogCount           int
	mu                 sync.RWMutex
	TriggerLogCount    int
	TriggerLogDuration time.Duration
	Enable             bool // If Enable is true, start the rule.
}

func NewLogCounter(logMode LogMode, triggerLogCount int, triggerLogDuration time.Duration) *LogCounter {
	logCounter := &LogCounter{}
	switch logMode {
	case AlwaysMode:
		logCounter.Enable = false
	case LowMode:
		logCounter.Enable = true
		logCounter.TriggerLogCount = 5
		logCounter.TriggerLogDuration = 60 * time.Second
	case MediumMode:
		logCounter.Enable = true
		logCounter.TriggerLogCount = 5
		logCounter.TriggerLogDuration = 300 * time.Second
	case HighMode:
		logCounter.Enable = true
		logCounter.TriggerLogCount = 3
		logCounter.TriggerLogDuration = 300 * time.Second
	case ForbiddenMode:
		logCounter.Enable = true
		logCounter.TriggerLogCount = 0
	}
	if triggerLogCount > 0 {
		logCounter.Enable = true
		logCounter.TriggerLogCount = triggerLogCount
		logCounter.TriggerLogDuration = triggerLogDuration
	}
	return logCounter
}

func (r *LogCounter) CheckPrintLog() bool
func (r *LogCounter) CheckDiffTime(lastErrorTime, newErrorTime time.Time) bool
