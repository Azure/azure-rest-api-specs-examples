Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a job step. This will implicitly create a new job version.
 *
 * @summary Creates or updates a job step. This will implicitly create a new job version.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CreateOrUpdateJobStepMax.json
 */
async function createOrUpdateAJobStepWithAllPropertiesSpecified() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "group1";
  const serverName = "server1";
  const jobAgentName = "agent1";
  const jobName = "job1";
  const stepName = "step1";
  const parameters = {
    action: { type: "TSql", source: "Inline", value: "select 2" },
    credential:
      "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/credentials/cred1",
    executionOptions: {
      initialRetryIntervalSeconds: 11,
      maximumRetryIntervalSeconds: 222,
      retryAttempts: 42,
      retryIntervalBackoffMultiplier: 3,
      timeoutSeconds: 1234,
    },
    output: {
      type: "SqlDatabase",
      credential:
        "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/credentials/cred0",
      databaseName: "database3",
      resourceGroupName: "group3",
      schemaName: "myschema1234",
      serverName: "server3",
      subscriptionId: "3501b905-a848-4b5d-96e8-b253f62d735a",
      tableName: "mytable5678",
    },
    stepId: 1,
    targetGroup:
      "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/servers/server1/jobAgents/agent1/targetGroups/targetGroup1",
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

createOrUpdateAJobStepWithAllPropertiesSpecified().catch(console.error);
```
