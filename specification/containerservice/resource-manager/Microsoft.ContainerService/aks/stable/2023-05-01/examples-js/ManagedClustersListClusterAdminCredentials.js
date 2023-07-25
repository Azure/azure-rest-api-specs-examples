const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the admin credentials of a managed cluster.
 *
 * @summary Lists the admin credentials of a managed cluster.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-05-01/examples/ManagedClustersListClusterAdminCredentials.json
 */
async function getManagedCluster() {
  const subscriptionId = process.env["CONTAINERSERVICE_SUBSCRIPTION_ID"] || "subid1";
  const resourceGroupName = process.env["CONTAINERSERVICE_RESOURCE_GROUP"] || "rg1";
  const resourceName = "clustername1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.managedClusters.listClusterAdminCredentials(
    resourceGroupName,
    resourceName
  );
  console.log(result);
}
