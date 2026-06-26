const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a job step. This will implicitly create a new job version.
 *
 * @summary creates or updates a job step. This will implicitly create a new job version.
 * x-ms-original-file: 2025-01-01/CreateOrUpdateJobStepMin.json
 */
async function createOrUpdateAJobStepWithMinimalPropertiesSpecified() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.jobSteps.createOrUpdate(
    "group1",
    "server1",
    "agent1",
    "job1",
    "step1",
    {
      action: { value: "select 1" },
      targetGroup:
        "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/targetGroups/targetGroup0",
    },
  );
  console.log(result);
}
