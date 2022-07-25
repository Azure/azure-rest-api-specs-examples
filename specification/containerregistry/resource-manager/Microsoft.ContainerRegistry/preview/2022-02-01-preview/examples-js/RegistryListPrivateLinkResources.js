const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the private link resources for a container registry.
 *
 * @summary Lists the private link resources for a container registry.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2022-02-01-preview/examples/RegistryListPrivateLinkResources.json
 */
async function registryListPrivateLinkResources() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const registryName = "myRegistry";
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.registries.listPrivateLinkResources(
    resourceGroupName,
    registryName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

registryListPrivateLinkResources().catch(console.error);
