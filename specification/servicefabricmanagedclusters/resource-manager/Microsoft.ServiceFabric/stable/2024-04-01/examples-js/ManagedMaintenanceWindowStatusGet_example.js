const {
  ServiceFabricManagedClustersManagementClient,
} = require("@azure/arm-servicefabricmanagedclusters");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Action to get Maintenance Window Status of the Service Fabric Managed Clusters.
 *
 * @summary Action to get Maintenance Window Status of the Service Fabric Managed Clusters.
 * x-ms-original-file: specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/stable/2024-04-01/examples/ManagedMaintenanceWindowStatusGet_example.json
 */
async function maintenanceWindowStatus() {
  const subscriptionId =
    process.env["SERVICEFABRICMANAGEDCLUSTERS_SUBSCRIPTION_ID"] ||
    "00000000-0000-0000-0000-000000000000";
  const resourceGroupName =
    process.env["SERVICEFABRICMANAGEDCLUSTERS_RESOURCE_GROUP"] || "resourceGroup1";
  const clusterName = "mycluster1";
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagedClustersManagementClient(credential, subscriptionId);
  const result = await client.managedMaintenanceWindowStatusOperations.get(
    resourceGroupName,
    clusterName,
  );
  console.log(result);
}
