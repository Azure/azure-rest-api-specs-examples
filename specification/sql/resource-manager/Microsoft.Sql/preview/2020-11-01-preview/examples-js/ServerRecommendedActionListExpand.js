const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of server advisors.
 *
 * @summary Gets a list of server advisors.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ServerRecommendedActionListExpand.json
 */
async function listOfServerRecommendedActionsForAllAdvisors() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "workloadinsight-demos";
  const serverName = "misosisvr";
  const expand = "recommendedActions";
  const options = { expand };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.serverAdvisors.listByServer(resourceGroupName, serverName, options);
  console.log(result);
}

listOfServerRecommendedActionsForAllAdvisors().catch(console.error);
