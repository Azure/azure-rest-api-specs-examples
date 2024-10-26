const { IoTOperationsClient } = require("@azure/arm-iotoperations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a DataflowEndpointResource
 *
 * @summary create a DataflowEndpointResource
 * x-ms-original-file: 2024-09-15-preview/DataflowEndpoint_CreateOrUpdate_ADX.json
 */
async function dataflowEndpointCreateOrUpdateAdx() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "F8C729F9-DF9C-4743-848F-96EE433D8E53";
  const client = new IoTOperationsClient(credential, subscriptionId);
  const result = await client.dataflowEndpoint.createOrUpdate(
    "rgiotoperations",
    "resource-name123",
    "adx-endpoint",
    {
      properties: {
        endpointType: "DataExplorer",
        dataExplorerSettings: {
          host: "example.westeurope.kusto.windows.net",
          authentication: {
            method: "SystemAssignedManagedIdentity",
            systemAssignedManagedIdentitySettings: {},
          },
          database: "example-database",
          batching: { latencySeconds: 9312, maxMessages: 9028 },
        },
      },
      extendedLocation: {
        name: "qmbrfwcpwwhggszhrdjv",
        type: "CustomLocation",
      },
    },
  );
  console.log(result);
}
