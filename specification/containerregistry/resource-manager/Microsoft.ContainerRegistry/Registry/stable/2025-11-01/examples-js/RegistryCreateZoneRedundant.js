const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a container registry with the specified parameters.
 *
 * @summary creates a container registry with the specified parameters.
 * x-ms-original-file: 2025-11-01/RegistryCreateZoneRedundant.json
 */
async function registryCreateZoneRedundant() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.registries.create("myResourceGroup", "myRegistry", {
    location: "westus",
    tags: { key: "value" },
    sku: { name: "Standard" },
    zoneRedundancy: "Enabled",
  });
  console.log(result);
}
