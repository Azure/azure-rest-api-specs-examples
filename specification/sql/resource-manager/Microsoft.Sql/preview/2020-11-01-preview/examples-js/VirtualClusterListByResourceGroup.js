const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of virtual clusters in a resource group.
 *
 * @summary Gets a list of virtual clusters in a resource group.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/VirtualClusterListByResourceGroup.json
 */
async function listVirtualClustersByResourceGroup() {
  const subscriptionId = "20d7082a-0fc7-4468-82bd-542694d5042b";
  const resourceGroupName = "testrg";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualClusters.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listVirtualClustersByResourceGroup().catch(console.error);
