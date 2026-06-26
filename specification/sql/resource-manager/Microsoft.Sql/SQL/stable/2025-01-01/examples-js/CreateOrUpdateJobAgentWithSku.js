const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a job agent.
 *
 * @summary creates or updates a job agent.
 * x-ms-original-file: 2025-01-01/CreateOrUpdateJobAgentWithSku.json
 */
async function createOrUpdateAJobAgentWithSku() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.jobAgents.createOrUpdate("group1", "server1", "agent1", {
    location: "southeastasia",
    databaseId:
      "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/databases/db1",
    sku: { name: "JA400" },
  });
  console.log(result);
}
