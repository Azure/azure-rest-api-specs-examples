const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the properties of the specified scope map.
 *
 * @summary Gets the properties of the specified scope map.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2022-02-01-preview/examples/ScopeMapGet.json
 */
async function scopeMapGet() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const registryName = "myRegistry";
  const scopeMapName = "myScopeMap";
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.scopeMaps.get(resourceGroupName, registryName, scopeMapName);
  console.log(result);
}

scopeMapGet().catch(console.error);
