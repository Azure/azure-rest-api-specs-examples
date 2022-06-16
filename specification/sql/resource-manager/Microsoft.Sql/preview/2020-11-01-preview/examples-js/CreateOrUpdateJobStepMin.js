const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a job step. This will implicitly create a new job version.
 *
 * @summary Creates or updates a job step. This will implicitly create a new job version.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CreateOrUpdateJobStepMin.json
 */
async function createOrUpdateAJobStepWithMinimalPropertiesSpecified() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "group1";
  const serverName = "server1";
  const jobAgentName = "agent1";
  const jobName = "job1";
  const stepName = "step1";
  const parameters = {
    action: { value: "select 1" },
    credential:
      "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/credentials/cred0",
    targetGroup:
      "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/targetGroups/targetGroup0",
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.jobSteps.createOrUpdate(
    resourceGroupName,
    serverName,
    jobAgentName,
    jobName,
    stepName,
    parameters
  );
  console.log(result);
}

createOrUpdateAJobStepWithMinimalPropertiesSpecified().catch(console.error);
