const { IoTOperationsClient } = require("@azure/arm-iotoperations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a DataflowEndpointResource
 *
 * @summary create a DataflowEndpointResource
 * x-ms-original-file: 2024-09-15-preview/DataflowEndpoint_CreateOrUpdate_EventGrid.json
 */
async function dataflowEndpointCreateOrUpdateEventGrid() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "F8C729F9-DF9C-4743-848F-96EE433D8E53";
  const client = new IoTOperationsClient(credential, subscriptionId);
  const result = await client.dataflowEndpoint.createOrUpdate(
    "rgiotoperations",
    "resource-name123",
    "event-grid-endpoint",
    {
      properties: {
        endpointType: "Mqtt",
        mqttSettings: {
          host: "example.westeurope-1.ts.eventgrid.azure.net:8883",
          authentication: {
            method: "SystemAssignedManagedIdentity",
            systemAssignedManagedIdentitySettings: {},
          },
          tls: { mode: "Enabled" },
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
