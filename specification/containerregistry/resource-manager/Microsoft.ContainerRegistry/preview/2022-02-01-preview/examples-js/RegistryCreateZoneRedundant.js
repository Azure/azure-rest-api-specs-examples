const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a container registry with the specified parameters.
 *
 * @summary Creates a container registry with the specified parameters.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2022-02-01-preview/examples/RegistryCreateZoneRedundant.json
 */
async function registryCreateZoneRedundant() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const registryName = "myRegistry";
  const registry = {
    location: "westus",
    sku: { name: "Standard" },
    tags: { key: "value" },
    zoneRedundancy: "Enabled",
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.registries.beginCreateAndWait(
    resourceGroupName,
    registryName,
    registry
  );
  console.log(result);
}

registryCreateZoneRedundant().catch(console.error);
