const { MySQLManagementClient } = require("@azure/arm-mysql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve recommended actions from the advisor.
 *
 * @summary Retrieve recommended actions from the advisor.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/RecommendedActionsGet.json
 */
async function recommendedActionsGet() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "testResourceGroupName";
  const serverName = "testServerName";
  const advisorName = "Index";
  const recommendedActionName = "Index-1";
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementClient(credential, subscriptionId);
  const result = await client.recommendedActions.get(
    resourceGroupName,
    serverName,
    advisorName,
    recommendedActionName
  );
  console.log(result);
}

recommendedActionsGet().catch(console.error);
