const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all job steps in the specified job version.
 *
 * @summary Gets all job steps in the specified job version.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ListJobStepsByVersion.json
 */
async function listJobStepsForTheSpecifiedVersionOfAJob() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "group1";
  const serverName = "server1";
  const jobAgentName = "agent1";
  const jobName = "job1";
  const jobVersion = 1;
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.jobSteps.listByVersion(
    resourceGroupName,
    serverName,
    jobAgentName,
    jobName,
    jobVersion
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listJobStepsForTheSpecifiedVersionOfAJob().catch(console.error);
