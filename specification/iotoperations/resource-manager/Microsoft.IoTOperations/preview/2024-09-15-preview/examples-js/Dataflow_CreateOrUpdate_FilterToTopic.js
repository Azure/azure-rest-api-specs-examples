const { IoTOperationsClient } = require("@azure/arm-iotoperations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a DataflowResource
 *
 * @summary create a DataflowResource
 * x-ms-original-file: 2024-09-15-preview/Dataflow_CreateOrUpdate_FilterToTopic.json
 */
async function dataflowCreateOrUpdateFilterToTopic() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "F8C729F9-DF9C-4743-848F-96EE433D8E53";
  const client = new IoTOperationsClient(credential, subscriptionId);
  const result = await client.dataflow.createOrUpdate(
    "rgiotoperations",
    "resource-name123",
    "resource-name123",
    "mqtt-filter-to-topic",
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
            name: "transformation1",
            builtInTransformationSettings: {
              filter: [
                {
                  type: "Filter",
                  description: "filter-datapoint",
                  inputs: ["temperature.Value", '"Tag 10".Value'],
                  expression: "$1 > 9000 && $2 >= 8000",
                },
              ],
              map: [{ type: "PassThrough", inputs: ["*"], output: "*" }],
            },
          },
          {
            operationType: "Destination",
            name: "destination1",
            destinationSettings: {
              endpointRef: "aio-builtin-broker-endpoint",
              dataDestination: "data/filtered/thermostat",
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
