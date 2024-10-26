const { IoTOperationsClient } = require("@azure/arm-iotoperations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a DataflowEndpointResource
 *
 * @summary create a DataflowEndpointResource
 * x-ms-original-file: 2024-09-15-preview/DataflowEndpoint_CreateOrUpdate_Fabric.json
 */
async function dataflowEndpointCreateOrUpdateFabric() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "F8C729F9-DF9C-4743-848F-96EE433D8E53";
  const client = new IoTOperationsClient(credential, subscriptionId);
  const result = await client.dataflowEndpoint.createOrUpdate(
    "rgiotoperations",
    "resource-name123",
    "fabric-endpoint",
    {
      properties: {
        endpointType: "FabricOneLake",
        fabricOneLakeSettings: {
          host: "onelake.dfs.fabric.microsoft.com",
          authentication: {
            method: "SystemAssignedManagedIdentity",
            systemAssignedManagedIdentitySettings: {},
          },
          names: {
            workspaceName: "example-workspace",
            lakehouseName: "example-lakehouse",
          },
          oneLakePathType: "Tables",
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
