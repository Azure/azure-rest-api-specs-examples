const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a connected registry for a container registry with the specified parameters.
 *
 * @summary Creates a connected registry for a container registry with the specified parameters.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2022-02-01-preview/examples/ConnectedRegistryCreate.json
 */
async function connectedRegistryCreate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const registryName = "myRegistry";
  const connectedRegistryName = "myConnectedRegistry";
  const connectedRegistryCreateParameters = {
    clientTokenIds: [
      "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/client1Token",
    ],
    mode: "ReadWrite",
    notificationsList: ["hello-world:*:*", "sample/repo/*:1.0:*"],
    parent: {
      syncProperties: {
        messageTtl: "P2D",
        schedule: "0 9 * * *",
        syncWindow: "PT3H",
        tokenId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/tokens/syncToken",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.connectedRegistries.beginCreateAndWait(
    resourceGroupName,
    registryName,
    connectedRegistryName,
    connectedRegistryCreateParameters
  );
  console.log(result);
}

connectedRegistryCreate().catch(console.error);
