// Copyright 2019 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package traceutil implements tracing utilities using "context".
package traceutil

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
)

const (
	TraceKey     = "trace"
	StartTimeKey = "startTime"
)

// Field is a kv pair to record additional details of the trace.
type Field struct {
	Key   string
	Value interface{}
}

func (f *Field) format() string {
	return fmt.Sprintf("%s:%v; ", f.Key, f.Value)
}

type Trace struct {
	operation    string
	lg           *zap.Logger
	fields       []Field
	startTime    time.Time
	steps        []step
	stepDisabled bool
	isEmpty      bool
}

type step struct {
	time            time.Time
	msg             string
	fields          []Field
	isSubTraceStart bool
	isSubTraceEnd   bool
}

func New(op string, lg *zap.Logger, fields ...Field) *Trace {
	return &Trace{operation: op, lg: lg, startTime: time.Now(), fields: fields}
}

// TODO returns a non-nil, empty Trace
func TODO() *Trace {
	return &Trace{isEmpty: true}
}

func Get(ctx context.Context) *Trace {
	if trace, ok := ctx.Value(TraceKey).(*Trace); ok && trace != nil {
		return trace
	}
	return TODO()
}

func (t *Trace) GetStartTime() time.Time {
	return t.startTime
}

func (t *Trace) SetStartTime(time time.Time) {
	t.startTime = time
}

func (t *Trace) InsertStep(at int, time time.Time, msg string, fields ...Field) {
	newStep := step{time: time, msg: msg, fields: fields}
	if at < len(t.steps) {
		t.steps = append(t.steps[:at+1], t.steps[at:]...)
		t.steps[at] = newStep
	} else {
		t.steps = append(t.steps, newStep)
	}
}

// StartSubTrace adds step to trace as a start sign of sublevel trace
// All steps in the subtrace will log out the input fields of this function
func (t *Trace) StartSubTrace(fields ...Field) {
	t.steps = append(t.steps, step{fields: fields, isSubTraceStart: true})
}

// StopSubTrace adds step to trace as a end sign of sublevel trace
// All steps in the subtrace will log out the input fields of this function
func (t *Trace) StopSubTrace(fields ...Field) {
	t.steps = append(t.steps, step{fields: fields, isSubTraceEnd: true})
}

// Step adds step to trace
func (t *Trace) Step(msg string, fields ...Field) {
	if !t.stepDisabled {
		t.steps = append(t.steps, step{time: time.Now(), msg: msg, fields: fields})
	}
}

// StepWithFunction 将测量输入函数作为一个单一的步骤
func (t *Trace) StepWithFunction(f func(), msg string, fields ...Field) {
	t.disableStep()
	f()
	t.enableStep()
	t.Step(msg, fields...)
}

func (t *Trace) AddField(fields ...Field) {
	for _, f := range fields {
		if !t.updateFieldIfExist(f) {
			t.fields = append(t.fields, f)
		}
	}
}

func (t *Trace) IsEmpty() bool {
	return t.isEmpty
}

func (t *Trace) updateFieldIfExist(f Field) bool {
	for i, v := range t.fields {
		if v.Key == f.Key {
			t.fields[i].Value = f.Value
			return true
		}
	}
	return false
}

// disableStep sets the flag to prevent the trace from adding steps
func (t *Trace) disableStep() {
	t.stepDisabled = true
}

// enableStep re-enable the trace to add steps
func (t *Trace) enableStep() {
	t.stepDisabled = false
}
