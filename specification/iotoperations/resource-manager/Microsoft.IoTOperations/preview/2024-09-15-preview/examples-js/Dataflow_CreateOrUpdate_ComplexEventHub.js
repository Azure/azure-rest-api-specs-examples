const { IoTOperationsClient } = require("@azure/arm-iotoperations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a DataflowResource
 *
 * @summary create a DataflowResource
 * x-ms-original-file: 2024-09-15-preview/Dataflow_CreateOrUpdate_ComplexEventHub.json
 */
async function dataflowCreateOrUpdateComplexEventHub() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "F8C729F9-DF9C-4743-848F-96EE433D8E53";
  const client = new IoTOperationsClient(credential, subscriptionId);
  const result = await client.dataflow.createOrUpdate(
    "rgiotoperations",
    "resource-name123",
    "resource-name123",
    "aio-to-event-hub-transformed",
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
              filter: [
                {
                  inputs: ["temperature.Value", '"Tag 10".Value'],
                  expression: "$1 > 9000 && $2 >= 8000",
                },
              ],
              map: [
                { inputs: ["*"], output: "*" },
                {
                  inputs: ["temperature.Value", '"Tag 10".Value'],
                  expression: "($1+$2)/2",
                  output: "AvgTemp.Value",
                },
                {
                  inputs: [],
                  expression: "true",
                  output: "dataflow-processed",
                },
                {
                  inputs: ["temperature.SourceTimestamp"],
                  expression: "",
                  output: "",
                },
                { inputs: ['"Tag 10"'], expression: "", output: "pressure" },
                {
                  inputs: ["temperature.Value"],
                  expression: "cToF($1)",
                  output: "temperatureF.Value",
                },
                {
                  inputs: ['"Tag 10".Value'],
                  expression: "scale ($1,0,10,0,100)",
                  output: '"Scale Tag 10".Value',
                },
              ],
            },
          },
          {
            operationType: "Destination",
            name: "destination1",
            destinationSettings: {
              endpointRef: "event-hub-endpoint",
              dataDestination: "myuniqueeventhub",
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
