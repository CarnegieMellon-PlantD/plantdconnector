# PlantD Connector

## Getting Started

1. Install OpenTelemetry builder.

   ```shell
   GO111MODULE=on go install go.opentelemetry.io/collector/cmd/builder@latest
   ```

   For more information about the `builder`, please refer to [OpenTelemetry Collector Builder README](https://github.com/open-telemetry/opentelemetry-collector/tree/main/cmd/builder#opentelemetry-collector-builder-ocb).

   Make sure the version of the builder is aligned with the version in the `builder-config.yaml`.

2. Build an executable collector binary.

   ```
   builder --config=builder-config.yaml
   ```

3. Run collector with a specific configuration.

   ```
   ./plantdcol
   ```

   