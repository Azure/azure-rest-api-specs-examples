const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a job.
 *
 * @summary creates or updates a job.
 * x-ms-original-file: 2025-01-01/CreateOrUpdateJobMin.json
 */
async function createAJobWithDefaultProperties() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new SqlManagementClient(credential, subscriptionId);
  const result = await client.jobs.createOrUpdate("group1", "server1", "agent1", "job1", {});
  console.log(result);
}
