package test

import (
	"fmt"
	"github.com/liuzhaomax/ovo/internal/core"
	"testing"
)

func TestTraceID(t *testing.T) {
	fmt.Println(core.TraceID())
	fmt.Println(core.SpanID())
}
