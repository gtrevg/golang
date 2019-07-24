// generated by jsonenums -type TaskState; DO NOT EDIT

package resillience

import (
	"encoding/json"
	"fmt"
)

var (
	_TaskStateNameToValue = map[string]TaskState{
		"TaskStateNew":               TaskStateNew,
		"TaskStateRunning":           TaskStateRunning,
		"TaskStateDoneErrorHappened": TaskStateDoneErrorHappened,
		"TaskStateDoneNormally":      TaskStateDoneNormally,
		"TaskStateDormancy":          TaskStateDormancy,
		"TaskStateDeath":             TaskStateDeath,
		"TaskStateButt":              TaskStateButt,
	}

	_TaskStateValueToName = map[TaskState]string{
		TaskStateNew:               "TaskStateNew",
		TaskStateRunning:           "TaskStateRunning",
		TaskStateDoneErrorHappened: "TaskStateDoneErrorHappened",
		TaskStateDoneNormally:      "TaskStateDoneNormally",
		TaskStateDormancy:          "TaskStateDormancy",
		TaskStateDeath:             "TaskStateDeath",
		TaskStateButt:              "TaskStateButt",
	}
)

func init() {
	var v TaskState
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_TaskStateNameToValue = map[string]TaskState{
			interface{}(TaskStateNew).(fmt.Stringer).String():               TaskStateNew,
			interface{}(TaskStateRunning).(fmt.Stringer).String():           TaskStateRunning,
			interface{}(TaskStateDoneErrorHappened).(fmt.Stringer).String(): TaskStateDoneErrorHappened,
			interface{}(TaskStateDoneNormally).(fmt.Stringer).String():      TaskStateDoneNormally,
			interface{}(TaskStateDormancy).(fmt.Stringer).String():          TaskStateDormancy,
			interface{}(TaskStateDeath).(fmt.Stringer).String():             TaskStateDeath,
			interface{}(TaskStateButt).(fmt.Stringer).String():              TaskStateButt,
		}
	}
}

// MarshalJSON is generated so TaskState satisfies json.Marshaler.
func (r TaskState) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _TaskStateValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid TaskState: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so TaskState satisfies json.Unmarshaler.
func (r *TaskState) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("TaskState should be a string, got %s", data)
	}
	v, ok := _TaskStateNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid TaskState %q", s)
	}
	*r = v
	return nil
}
