const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a target execution.
 *
 * @summary Gets a target execution.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/GetJobExecutionTarget.json
 */
async function getAJobStepTargetExecution() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "group1";
  const serverName = "server1";
  const jobAgentName = "agent1";
  const jobName = "job1";
  const jobExecutionId = "5555-6666-7777-8888-999999999999";
  const stepName = "step1";
  const targetId = "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.jobTargetExecutions.get(
    resourceGroupName,
    serverName,
    jobAgentName,
    jobName,
    jobExecutionId,
    stepName,
    targetId
  );
  console.log(result);
}

getAJobStepTargetExecution().catch(console.error);
