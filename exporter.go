package consoleexporter

import (
    "context"
    "fmt"

    "go.opentelemetry.io/collector/pdata/pcommon"
    "go.opentelemetry.io/collector/pdata/ptrace"
)

// consoleExporter just implements a very basic traces exporter.
type consoleExporter struct{}

// ConsumeTraces is called with each batch of incoming traces.
func (e *consoleExporter) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
    rs := td.ResourceSpans()
    for i := 0; i < rs.Len(); i++ {
        resource := rs.At(i).Resource()
        // Print resource attributes (e.g. service.name)
        resource.Attributes().Range(func(k string, v pcommon.Value) bool {
            fmt.Printf("Resource %q = %v\n", k, v.AsString())
            return true
        })

        ilss := rs.At(i).ScopeSpans()
        for j := 0; j < ilss.Len(); j++ {
            spans := ilss.At(j).Spans()
            for k := 0; k < spans.Len(); k++ {
                span := spans.At(k)
                fmt.Printf(
                    "TraceID=%s  SpanName=%s  Start=%s  End=%s\n",
                    span.TraceID().HexString(),
                    span.Name(),
                    span.StartTimestamp().AsTime().Format("15:04:05.000"),
                    span.EndTimestamp().AsTime().Format("15:04:05.000"),
                )
            }
        }
    }
    return nil
}

// Start is a no-op.
func (e *consoleExporter) Start(ctx context.Context, host component.Host) error {
    fmt.Println("consoleexporter starting")
    return nil
}

// Shutdown is a no-op.
func (e *consoleExporter) Shutdown(ctx context.Context) error {
    fmt.Println("consoleexporter shutting down")
    return nil
}
