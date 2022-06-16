const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified version of a job step.
 *
 * @summary Gets the specified version of a job step.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/GetJobStepByVersion.json
 */
async function getTheSpecifiedVersionOfAJobStep() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "group1";
  const serverName = "server1";
  const jobAgentName = "agent1";
  const jobName = "job1";
  const jobVersion = 1;
  const stepName = "step1";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.jobSteps.getByVersion(
    resourceGroupName,
    serverName,
    jobAgentName,
    jobName,
    jobVersion,
    stepName
  );
  console.log(result);
}

getTheSpecifiedVersionOfAJobStep().catch(console.error);
