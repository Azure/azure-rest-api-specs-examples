const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a container registry with the specified parameters.
 *
 * @summary creates a container registry with the specified parameters.
 * x-ms-original-file: 2026-11-01/RegistryCreate.json
 */
async function registryCreate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.registries.create("myResourceGroup", "myRegistry", {
    location: "westus",
    tags: { key: "value" },
    sku: { name: "Premium" },
    adminUserEnabled: true,
    regionalEndpoints: "Enabled",
    dataEndpointEnabled: true,
    endpointProtocol: "IPv4AndIPv6",
  });
  console.log(result);
}
