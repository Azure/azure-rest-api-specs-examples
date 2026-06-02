const { DataBoxManagementClient } = require("@azure/arm-databox");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets information about the specified job.
 *
 * @summary gets information about the specified job.
 * x-ms-original-file: 2025-07-01/JobsGetExport.json
 */
async function jobsGetExport() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "YourSubscriptionId";
  const client = new DataBoxManagementClient(credential, subscriptionId);
  const result = await client.jobs.get("YourResourceGroupName", "TestJobName1", {
    expand: "details",
  });
  console.log(result);
}
