const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of server advisors.
 *
 * @summary Gets a list of server advisors.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ServerAdvisorList.json
 */
async function listOfServerAdvisors() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "workloadinsight-demos";
  const serverName = "misosisvr";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.serverAdvisors.listByServer(resourceGroupName, serverName);
  console.log(result);
}
