const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deactivates the connected registry instance.
 *
 * @summary Deactivates the connected registry instance.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2022-02-01-preview/examples/ConnectedRegistryDeactivate.json
 */
async function connectedRegistryDeactivate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const registryName = "myRegistry";
  const connectedRegistryName = "myConnectedRegistry";
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.connectedRegistries.beginDeactivateAndWait(
    resourceGroupName,
    registryName,
    connectedRegistryName
  );
  console.log(result);
}

connectedRegistryDeactivate().catch(console.error);
