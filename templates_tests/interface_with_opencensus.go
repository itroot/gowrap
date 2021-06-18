package templatestests

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using ../templates/opencensus template

//go:generate gowrap gen -p github.com/hexdigest/gowrap/templates_tests -i TestInterface -t ../templates/opencensus -o interface_with_opencensus.go -v DecoratorName=TestInterfaceWithOpenCensus

import (
	"context"

	"go.opencensus.io/trace"
)

// TestInterfaceWithOpenCensus implements TestInterface interface instrumented with opentracing spans
type TestInterfaceWithOpenCensus struct {
	TestInterface
	_instance      string
	_spanDecorator func(span *trace.Span, params, results map[string]interface{})
}

// NewTestInterfaceWithOpenCensus returns TestInterfaceWithOpenCensus
func NewTestInterfaceWithOpenCensus(base TestInterface, instance string, spanDecorator ...func(span *trace.Span, params, results map[string]interface{})) TestInterfaceWithOpenCensus {
	d := TestInterfaceWithOpenCensus{
		TestInterface: base,
		_instance:     instance,
	}

	if len(spanDecorator) > 0 && spanDecorator[0] != nil {
		d._spanDecorator = spanDecorator[0]
	}

	return d
}

// F implements TestInterface
func (_d TestInterfaceWithOpenCensus) F(ctx context.Context, a1 string, a2 ...string) (result1 string, result2 string, err error) {
	ctx, _span := trace.StartSpan(ctx, _d._instance+".TestInterface.F")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"a1":  a1,
				"a2":  a2}, map[string]interface{}{
				"result1": result1,
				"result2": result2,
				"err":     err})
		} else if err != nil {
			_span.AddAttributes(
				trace.BoolAttribute("error", true),
				trace.StringAttribute("event", "error"),
				trace.StringAttribute("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.TestInterface.F(ctx, a1, a2...)
}
