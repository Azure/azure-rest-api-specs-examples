const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Requests cancellation of a job execution.
 *
 * @summary Requests cancellation of a job execution.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CancelJobExecution.json
 */
async function cancelAJobExecution() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "group1";
  const serverName = "server1";
  const jobAgentName = "agent1";
  const jobName = "job1";
  const jobExecutionId = "5555-6666-7777-8888-999999999999";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.jobExecutions.cancel(
    resourceGroupName,
    serverName,
    jobAgentName,
    jobName,
    jobExecutionId
  );
  console.log(result);
}

cancelAJobExecution().catch(console.error);
