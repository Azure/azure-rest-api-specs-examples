const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a target execution.
 *
 * @summary Gets a target execution.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/GetJobExecutionTarget.json
 */
async function getAJobStepTargetExecution() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "group1";
  const serverName = "server1";
  const jobAgentName = "agent1";
  const jobName = "job1";
  const jobExecutionId = "5A86BF65-43AC-F258-2524-9E92992F97CA";
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
    targetId,
  );
  console.log(result);
}
