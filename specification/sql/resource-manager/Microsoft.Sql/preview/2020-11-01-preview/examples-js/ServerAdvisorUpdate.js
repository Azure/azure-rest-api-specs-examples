const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a server advisor.
 *
 * @summary Updates a server advisor.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ServerAdvisorUpdate.json
 */
async function updateServerAdvisor() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "workloadinsight-demos";
  const serverName = "misosisvr";
  const advisorName = "CreateIndex";
  const parameters = { autoExecuteStatus: "Disabled" };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.serverAdvisors.update(
    resourceGroupName,
    serverName,
    advisorName,
    parameters
  );
  console.log(result);
}

updateServerAdvisor().catch(console.error);
