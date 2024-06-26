const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a job.
 *
 * @summary Creates or updates a job.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CreateOrUpdateJobMax.json
 */
async function createAJobWithAllPropertiesSpecified() {
  const subscriptionId =
    process.env["SQL_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SQL_RESOURCE_GROUP"] || "group1";
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
    parameters,
  );
  console.log(result);
}
