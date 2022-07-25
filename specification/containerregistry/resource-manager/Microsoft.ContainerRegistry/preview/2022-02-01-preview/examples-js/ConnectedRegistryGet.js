const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the properties of the connected registry.
 *
 * @summary Gets the properties of the connected registry.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2022-02-01-preview/examples/ConnectedRegistryGet.json
 */
async function connectedRegistryGet() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const registryName = "myRegistry";
  const connectedRegistryName = "myConnectedRegistry";
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.connectedRegistries.get(
    resourceGroupName,
    registryName,
    connectedRegistryName
  );
  console.log(result);
}

connectedRegistryGet().catch(console.error);
