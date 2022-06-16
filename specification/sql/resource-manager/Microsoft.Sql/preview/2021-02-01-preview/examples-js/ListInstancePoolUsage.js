const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all instance pool usage metrics
 *
 * @summary Gets all instance pool usage metrics
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/ListInstancePoolUsage.json
 */
async function listInstancePoolUsages() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "group1";
  const instancePoolName = "testIP";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.usages.listByInstancePool(resourceGroupName, instancePoolName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listInstancePoolUsages().catch(console.error);
