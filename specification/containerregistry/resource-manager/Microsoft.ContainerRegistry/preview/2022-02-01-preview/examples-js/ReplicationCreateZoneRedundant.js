const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a replication for a container registry with the specified parameters.
 *
 * @summary Creates a replication for a container registry with the specified parameters.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2022-02-01-preview/examples/ReplicationCreateZoneRedundant.json
 */
async function replicationCreateZoneRedundant() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const registryName = "myRegistry";
  const replicationName = "myReplication";
  const replication = {
    location: "eastus",
    regionEndpointEnabled: true,
    tags: { key: "value" },
    zoneRedundancy: "Enabled",
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.replications.beginCreateAndWait(
    resourceGroupName,
    registryName,
    replicationName,
    replication
  );
  console.log(result);
}

replicationCreateZoneRedundant().catch(console.error);
