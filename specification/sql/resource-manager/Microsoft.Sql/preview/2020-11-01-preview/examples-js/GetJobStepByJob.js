const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a job step in a job's current version.
 *
 * @summary Gets a job step in a job's current version.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/GetJobStepByJob.json
 */
async function getTheLatestVersionOfAJobStep() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "group1";
  const serverName = "server1";
  const jobAgentName = "agent1";
  const jobName = "job1";
  const stepName = "step1";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.jobSteps.get(
    resourceGroupName,
    serverName,
    jobAgentName,
    jobName,
    stepName
  );
  console.log(result);
}

getTheLatestVersionOfAJobStep().catch(console.error);
