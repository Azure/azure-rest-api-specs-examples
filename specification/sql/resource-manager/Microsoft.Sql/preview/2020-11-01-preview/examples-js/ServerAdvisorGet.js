const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a server advisor.
 *
 * @summary Gets a server advisor.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ServerAdvisorGet.json
 */
async function getServerAdvisor() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "workloadinsight-demos";
  const serverName = "misosisvr";
  const advisorName = "CreateIndex";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.serverAdvisors.get(resourceGroupName, serverName, advisorName);
  console.log(result);
}

getServerAdvisor().catch(console.error);
