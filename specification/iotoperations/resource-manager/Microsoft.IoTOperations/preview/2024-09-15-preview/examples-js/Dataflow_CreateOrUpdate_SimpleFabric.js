const { IoTOperationsClient } = require("@azure/arm-iotoperations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a DataflowResource
 *
 * @summary create a DataflowResource
 * x-ms-original-file: 2024-09-15-preview/Dataflow_CreateOrUpdate_SimpleFabric.json
 */
async function dataflowCreateOrUpdateSimpleFabric() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "F8C729F9-DF9C-4743-848F-96EE433D8E53";
  const client = new IoTOperationsClient(credential, subscriptionId);
  const result = await client.dataflow.createOrUpdate(
    "rgiotoperations",
    "resource-name123",
    "resource-name123",
    "aio-to-fabric",
    {
      properties: {
        mode: "Enabled",
        operations: [
          {
            operationType: "Source",
            name: "source1",
            sourceSettings: {
              endpointRef: "aio-builtin-broker-endpoint",
              dataSources: ["azure-iot-operations/data/thermostat"],
            },
          },
          {
            operationType: "BuiltInTransformation",
            builtInTransformationSettings: {
              serializationFormat: "Parquet",
              schemaRef: "aio-sr://exampleNamespace/exmapleParquetSchema:1.0.0",
            },
          },
          {
            operationType: "Destination",
            name: "destination1",
            destinationSettings: {
              endpointRef: "fabric-endpoint",
              dataDestination: "telemetryTable",
            },
          },
        ],
      },
      extendedLocation: {
        name: "qmbrfwcpwwhggszhrdjv",
        type: "CustomLocation",
      },
    },
  );
  console.log(result);
}
