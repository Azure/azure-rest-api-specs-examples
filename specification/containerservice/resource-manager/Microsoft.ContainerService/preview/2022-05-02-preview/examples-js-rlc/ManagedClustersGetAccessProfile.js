const ContainerServiceManagementClient = require("@azure-rest/arm-containerservice").default;
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to **WARNING**: This API will be deprecated. Instead use [ListClusterUserCredentials](https://learn.microsoft.com/rest/api/aks/managedclusters/listclusterusercredentials) or [ListClusterAdminCredentials](https://learn.microsoft.com/rest/api/aks/managedclusters/listclusteradmincredentials) .
 *
 * @summary **WARNING**: This API will be deprecated. Instead use [ListClusterUserCredentials](https://learn.microsoft.com/rest/api/aks/managedclusters/listclusterusercredentials) or [ListClusterAdminCredentials](https://learn.microsoft.com/rest/api/aks/managedclusters/listclusteradmincredentials) .
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-05-02-preview/examples/ManagedClustersGetAccessProfile.json
 */
async function getManagedCluster() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const roleName = "clusterUser";
  const credential = new DefaultAzureCredential();
  const client = ContainerServiceManagementClient(credential);
  const result = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/accessProfiles/{roleName}/listCredential",
      subscriptionId,
      resourceGroupName,
      resourceName,
      roleName,
    )
    .post();
  console.log(result);
}

getManagedCluster().catch(console.error);
