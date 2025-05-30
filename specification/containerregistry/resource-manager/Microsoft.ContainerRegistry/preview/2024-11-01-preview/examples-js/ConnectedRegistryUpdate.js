const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Updates a connected registry with the specified parameters.
 *
 * @summary Updates a connected registry with the specified parameters.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2024-11-01-preview/examples/ConnectedRegistryUpdate.json
 */
async function connectedRegistryUpdate() {
  const subscriptionId =
    process.env["CONTAINERREGISTRY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONTAINERREGISTRY_RESOURCE_GROUP"] || "myResourceGroup";
  const registryName = "myRegistry";
  const connectedRegistryName = "myScopeMap";
  const connectedRegistryUpdateParameters = {
    clientTokenIds: [
      "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client1Token",
      "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client2Token",
    ],
    garbageCollection: { enabled: true, schedule: "0 5 * * *" },
    logging: { auditLogStatus: "Enabled", logLevel: "Debug" },
    notificationsList: ["hello-world:*:*", "sample/repo/*:1.0:*"],
    syncProperties: {
      messageTtl: "P30D",
      schedule: "0 0 */10 * *",
      syncWindow: "P2D",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.connectedRegistries.beginUpdateAndWait(
    resourceGroupName,
    registryName,
    connectedRegistryName,
    connectedRegistryUpdateParameters,
  );
  console.log(result);
}
