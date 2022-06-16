const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the target executions of a job step execution.
 *
 * @summary Lists the target executions of a job step execution.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ListJobExecutionTargetsByStep.json
 */
async function listJobStepTargetExecutions() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "group1";
  const serverName = "server1";
  const jobAgentName = "agent1";
  const jobName = "job1";
  const jobExecutionId = "5555-6666-7777-8888-999999999999";
  const stepName = "step1";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.jobTargetExecutions.listByStep(
    resourceGroupName,
    serverName,
    jobAgentName,
    jobName,
    jobExecutionId,
    stepName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listJobStepTargetExecutions().catch(console.error);
