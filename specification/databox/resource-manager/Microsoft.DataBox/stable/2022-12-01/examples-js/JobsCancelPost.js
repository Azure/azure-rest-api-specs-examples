const { DataBoxManagementClient } = require("@azure/arm-databox");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to CancelJob.
 *
 * @summary CancelJob.
 * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2022-12-01/examples/JobsCancelPost.json
 */
async function jobsCancelPost() {
  const subscriptionId = process.env["DATABOX_SUBSCRIPTION_ID"] || "YourSubscriptionId";
  const resourceGroupName = process.env["DATABOX_RESOURCE_GROUP"] || "YourResourceGroupName";
  const jobName = "TestJobName1";
  const cancellationReason = { reason: "CancelTest" };
  const credential = new DefaultAzureCredential();
  const client = new DataBoxManagementClient(credential, subscriptionId);
  const result = await client.jobs.cancel(resourceGroupName, jobName, cancellationReason);
  console.log(result);
}
