const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists a job's executions.
 *
 * @summary Lists a job's executions.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ListJobExecutionsByJob.json
 */
async function listAJobExecutions() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "group1";
  const serverName = "server1";
  const jobAgentName = "agent1";
  const jobName = "job1";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.jobExecutions.listByJob(
    resourceGroupName,
    serverName,
    jobAgentName,
    jobName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAJobExecutions().catch(console.error);
