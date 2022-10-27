const ContainerServiceManagementClient = require("@azure-rest/arm-containerservice").default;
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a managed cluster.
 *
 * @summary Deletes a managed cluster.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-05-02-preview/examples/ManagedClustersDelete.json
 */
async function deleteManagedCluster() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const credential = new DefaultAzureCredential();
  const client = ContainerServiceManagementClient(credential);
  const result = client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}",
      subscriptionId,
      resourceGroupName,
      resourceName
    )
    .delete();
  console.log(result);
}

deleteManagedCluster().catch(console.error);
