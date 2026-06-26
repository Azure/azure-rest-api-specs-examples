const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates a job agent.
 *
 * @summary updates a job agent.
 * x-ms-original-file: 2025-01-01/UpdateJobAgentWithSku.json
 */
async function updateAJobAgentSku() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.jobAgents.update("group1", "server1", "agent1", {
    sku: { name: "JA200" },
  });
  console.log(result);
}
