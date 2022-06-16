const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a sync agent.
 *
 * @summary Deletes a sync agent.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/SyncAgentDelete.json
 */
async function deleteASyncAgent() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "syncagentcrud-65440";
  const serverName = "syncagentcrud-8475";
  const syncAgentName = "syncagentcrud-3187";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.syncAgents.beginDeleteAndWait(
    resourceGroupName,
    serverName,
    syncAgentName
  );
  console.log(result);
}

deleteASyncAgent().catch(console.error);
