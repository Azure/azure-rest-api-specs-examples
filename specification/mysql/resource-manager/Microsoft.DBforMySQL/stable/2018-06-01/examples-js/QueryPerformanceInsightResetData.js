const { MySQLManagementClient } = require("@azure/arm-mysql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Reset data for Query Performance Insight.
 *
 * @summary Reset data for Query Performance Insight.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/QueryPerformanceInsightResetData.json
 */
async function queryPerformanceInsightResetData() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "testResourceGroupName";
  const serverName = "testServerName";
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementClient(credential, subscriptionId);
  const result = await client.resetQueryPerformanceInsightData(resourceGroupName, serverName);
  console.log(result);
}

queryPerformanceInsightResetData().catch(console.error);
