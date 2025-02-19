const {
  ServiceFabricManagedClustersManagementClient,
} = require("@azure/arm-servicefabricmanagedclusters");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Send a request to start a rollback of the current application upgrade. This will start rolling back the application to the previous version.
 *
 * @summary Send a request to start a rollback of the current application upgrade. This will start rolling back the application to the previous version.
 * x-ms-original-file: specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-09-01-preview/examples/ApplicationActionStartRollback_example.json
 */
async function startAnApplicationUpgradeRollback() {
  const subscriptionId =
    process.env["SERVICEFABRICMANAGEDCLUSTERS_SUBSCRIPTION_ID"] ||
    "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SERVICEFABRICMANAGEDCLUSTERS_RESOURCE_GROUP"] || "resRg";
  const clusterName = "myCluster";
  const applicationName = "myApp";
  const credential = new DefaultAzureCredential();
  const client = new ServiceFabricManagedClustersManagementClient(credential, subscriptionId);
  const result = await client.applications.beginStartRollbackAndWait(
    resourceGroupName,
    clusterName,
    applicationName,
  );
  console.log(result);
}
