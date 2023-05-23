const { DataBoxManagementClient } = require("@azure/arm-databox");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a job.
 *
 * @summary Deletes a job.
 * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2022-12-01/examples/JobsDelete.json
 */
async function jobsDelete() {
  const subscriptionId = process.env["DATABOX_SUBSCRIPTION_ID"] || "YourSubscriptionId";
  const resourceGroupName = process.env["DATABOX_RESOURCE_GROUP"] || "YourResourceGroupName";
  const jobName = "TestJobName1";
  const credential = new DefaultAzureCredential();
  const client = new DataBoxManagementClient(credential, subscriptionId);
  const result = await client.jobs.beginDeleteAndWait(resourceGroupName, jobName);
  console.log(result);
}
