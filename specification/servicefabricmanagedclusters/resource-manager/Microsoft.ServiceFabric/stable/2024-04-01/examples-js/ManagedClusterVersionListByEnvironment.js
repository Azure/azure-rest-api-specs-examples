const {
  ServiceFabricManagedClustersManagementClient,
} = require("@azure/arm-servicefabricmanagedclusters");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all available code versions for Service Fabric cluster resources by environment.
 *
 * @summary Gets all available code versions for Service Fabric cluster resources by environment.
 * x-ms-original-file: specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/stable/2024-04-01/examples/ManagedClusterVersionListByEnvironment.json
 */
async function listClusterVersionsByEnvironment() {
  const subscriptionId =
    process.env["SERVICEFABRICMANAGEDCLUSTERS_SUBSCRIPTION_ID"] ||
    "00000000-0000-0000-0000-000000000000";
  const location = "eastus";
  const environment = "Windows";
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagedClustersManagementClient(credential, subscriptionId);
  const result = await client.managedClusterVersion.listByEnvironment(location, environment);
  console.log(result);
}
