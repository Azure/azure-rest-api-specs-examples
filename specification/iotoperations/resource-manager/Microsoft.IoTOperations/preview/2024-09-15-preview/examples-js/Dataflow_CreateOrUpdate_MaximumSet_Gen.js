const { IoTOperationsClient } = require("@azure/arm-iotoperations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a DataflowResource
 *
 * @summary create a DataflowResource
 * x-ms-original-file: 2024-09-15-preview/Dataflow_CreateOrUpdate_MaximumSet_Gen.json
 */
async function dataflowCreateOrUpdate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "F8C729F9-DF9C-4743-848F-96EE433D8E53";
  const client = new IoTOperationsClient(credential, subscriptionId);
  const result = await client.dataflow.createOrUpdate(
    "rgiotoperations",
    "resource-name123",
    "resource-name123",
    "resource-name123",
    {
      properties: {
        mode: "Enabled",
        operations: [
          {
            operationType: "Source",
            name: "knnafvkwoeakm",
            sourceSettings: {
              endpointRef: "iixotodhvhkkfcfyrkoveslqig",
              assetRef: "zayyykwmckaocywdkohmu",
              serializationFormat: "Json",
              schemaRef: "pknmdzqll",
              dataSources: ["chkkpymxhp"],
            },
            builtInTransformationSettings: {
              serializationFormat: "Delta",
              schemaRef: "mcdc",
              datasets: [
                {
                  key: "qsfqcgxaxnhfumrsdsokwyv",
                  description: "Lorem ipsum odor amet, consectetuer adipiscing elit.",
                  schemaRef: "n",
                  inputs: ["mosffpsslifkq"],
                  expression: "aatbwomvflemsxialv",
                },
              ],
              filter: [
                {
                  type: "Filter",
                  description: "Lorem ipsum odor amet, consectetuer adipiscing elit.",
                  inputs: ["sxmjkbntgb"],
                  expression: "n",
                },
              ],
              map: [
                {
                  type: "NewProperties",
                  description: "Lorem ipsum odor amet, consectetuer adipiscing elit.",
                  inputs: ["xsbxuk"],
                  expression: "txoiltogsarwkzalsphvlmt",
                  output: "nvgtmkfl",
                },
              ],
            },
            destinationSettings: {
              endpointRef: "kybkchnzimerguekuvqlqiqdvvrt",
              dataDestination: "cbrh",
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
