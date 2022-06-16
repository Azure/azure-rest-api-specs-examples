const { MySQLManagementClient } = require("@azure/arm-mysql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a recommendation action advisor.
 *
 * @summary Get a recommendation action advisor.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/AdvisorsGet.json
 */
async function advisorsGet() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "testResourceGroupName";
  const serverName = "testServerName";
  const advisorName = "Index";
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementClient(credential, subscriptionId);
  const result = await client.advisors.get(resourceGroupName, serverName, advisorName);
  console.log(result);
}

advisorsGet().catch(console.error);
