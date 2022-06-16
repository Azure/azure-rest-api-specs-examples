const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create cluster identity.
 *
 * @summary Create cluster identity.
 * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2022-05-01/examples/CreateClusterIdentity.json
 */
async function createClusterIdentity() {
  const subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = "test-rg";
  const clusterName = "myCluster";
  const credential = new DefaultAzureCredential();
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const result = await client.clusters.beginCreateIdentityAndWait(resourceGroupName, clusterName);
  console.log(result);
}

createClusterIdentity().catch(console.error);
