const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the step executions of a job execution.
 *
 * @summary Lists the step executions of a job execution.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ListJobExecutionSteps.json
 */
async function listJobStepExecutions() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "group1";
  const serverName = "server1";
  const jobAgentName = "agent1";
  const jobName = "job1";
  const jobExecutionId = "5555-6666-7777-8888-999999999999";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.jobStepExecutions.listByJobExecution(
    resourceGroupName,
    serverName,
    jobAgentName,
    jobName,
    jobExecutionId
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listJobStepExecutions().catch(console.error);
