const { MySQLManagementClient } = require("@azure/arm-mysql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve recommended actions from the advisor.
 *
 * @summary Retrieve recommended actions from the advisor.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/RecommendedActionsListByServer.json
 */
async function recommendedActionsListByServer() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "testResourceGroupName";
  const serverName = "testServerName";
  const advisorName = "Index";
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.recommendedActions.listByServer(
    resourceGroupName,
    serverName,
    advisorName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

recommendedActionsListByServer().catch(console.error);
