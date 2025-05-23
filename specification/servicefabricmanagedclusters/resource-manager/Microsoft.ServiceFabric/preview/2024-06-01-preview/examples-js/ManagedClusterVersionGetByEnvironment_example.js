const {
  ServiceFabricManagedClustersManagementClient,
} = require("@azure/arm-servicefabricmanagedclusters");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about an available Service Fabric cluster code version by environment.
 *
 * @summary Gets information about an available Service Fabric cluster code version by environment.
 * x-ms-original-file: specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-06-01-preview/examples/ManagedClusterVersionGetByEnvironment_example.json
 */
async function getClusterVersionByEnvironment() {
  const subscriptionId =
    process.env["SERVICEFABRICMANAGEDCLUSTERS_SUBSCRIPTION_ID"] ||
    "00000000-0000-0000-0000-000000000000";
  const location = "eastus";
  const environment = "Windows";
  const clusterVersion = "7.2.477.9590";
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagedClustersManagementClient(credential, subscriptionId);
  const result = await client.managedClusterVersion.getByEnvironment(
    location,
    environment,
    clusterVersion,
  );
  console.log(result);
}
