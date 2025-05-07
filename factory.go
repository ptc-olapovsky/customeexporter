package consoleexporter

import (
    "context"

    "go.opentelemetry.io/collector/component"
    "go.opentelemetry.io/collector/config/configmodels"
    "go.opentelemetry.io/collector/exporter/exporterhelper"
    "go.opentelemetry.io/collector/pdata/ptrace"
)

// Config defines exporter settings (none beyond the defaults here).
type Config struct {
    configmodels.ExporterSettings `mapstructure:",squash"`
}

// NewFactory returns a factory for the console exporter.
func NewFactory() component.ExporterFactory {
    return exporterhelper.NewFactory(
        "console",
        func() component.Config { return &Config{ExporterSettings: configmodels.ExporterSettings{TypeVal: "console", NameVal: "console"}} },
        exporterhelper.WithTraces(
            func(
                ctx context.Context,
                set exporterhelper.CreateSettings,
                td ptrace.Traces,
            ) error {
                return (&consoleExporter{}).ConsumeTraces(ctx, td)
            },
            exporterhelper.WithStart((&consoleExporter{}).Start),
            exporterhelper.WithShutdown((&consoleExporter{}).Shutdown),
        ),
    )
}
