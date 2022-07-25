const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a private link resource by a specified group name for a container registry.
 *
 * @summary Gets a private link resource by a specified group name for a container registry.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2022-02-01-preview/examples/RegistryGetPrivateLinkResource.json
 */
async function registryGetPrivateLinkResource() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const registryName = "myRegistry";
  const groupName = "registry";
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.registries.getPrivateLinkResource(
    resourceGroupName,
    registryName,
    groupName
  );
  console.log(result);
}

registryGetPrivateLinkResource().catch(console.error);
