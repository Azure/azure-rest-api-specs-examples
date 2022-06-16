const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a virtual cluster.
 *
 * @summary Gets a virtual cluster.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/VirtualClusterGet.json
 */
async function getVirtualCluster() {
  const subscriptionId = "20d7082a-0fc7-4468-82bd-542694d5042b";
  const resourceGroupName = "testrg";
  const virtualClusterName = "vc-subnet1-f769ed71-b3ad-491a-a9d5-26eeceaa6be2";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.virtualClusters.get(resourceGroupName, virtualClusterName);
  console.log(result);
}

getVirtualCluster().catch(console.error);
