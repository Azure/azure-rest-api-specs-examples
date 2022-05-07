Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-sql_9.0.1/sdk/sql/arm-sql/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a job.
 *
 * @summary Creates or updates a job.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CreateOrUpdateJobMax.json
 */
async function createAJobWithAllPropertiesSpecified() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "group1";
  const serverName = "server1";
  const jobAgentName = "agent1";
  const jobName = "job1";
  const parameters = {
    description: "my favourite job",
    schedule: {
      type: "Recurring",
      enabled: true,
      endTime: new Date("2015-09-24T23:59:59Z"),
      interval: "PT5M",
      startTime: new Date("2015-09-24T18:30:01Z"),
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.jobs.createOrUpdate(
    resourceGroupName,
    serverName,
    jobAgentName,
    jobName,
    parameters
  );
  console.log(result);
}

createAJobWithAllPropertiesSpecified().catch(console.error);
```
