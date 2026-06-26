const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists all executions in a job agent.
 *
 * @summary lists all executions in a job agent.
 * x-ms-original-file: 2025-01-01/ListJobExecutionsByAgent.json
 */
async function listAllJobExecutionsInAJobAgent() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.jobExecutions.listByAgent("group1", "server1", "agent1")) {
    resArray.push(item);
  }

  console.log(resArray);
}
