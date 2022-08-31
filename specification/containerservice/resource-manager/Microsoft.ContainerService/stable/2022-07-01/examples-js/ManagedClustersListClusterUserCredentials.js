const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the user credentials of a managed cluster.
 *
 * @summary Lists the user credentials of a managed cluster.
 * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-07-01/examples/ManagedClustersListClusterUserCredentials.json
 */
async function getManagedCluster() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.managedClusters.listClusterUserCredentials(
    resourceGroupName,
    resourceName
  );
  console.log(result);
}

getManagedCluster().catch(console.error);
