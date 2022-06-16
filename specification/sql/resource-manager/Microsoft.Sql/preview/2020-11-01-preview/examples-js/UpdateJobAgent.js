const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a job agent.
 *
 * @summary Updates a job agent.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/UpdateJobAgent.json
 */
async function updateAJobAgentTags() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "group1";
  const serverName = "server1";
  const jobAgentName = "agent1";
  const parameters = { tags: { mytag1: "myvalue1" } };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.jobAgents.beginUpdateAndWait(
    resourceGroupName,
    serverName,
    jobAgentName,
    parameters
  );
  console.log(result);
}

updateAJobAgentTags().catch(console.error);
