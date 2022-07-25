const { DataBoxManagementClient } = require("@azure/arm-databox");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to CancelJob.
 *
 * @summary CancelJob.
 * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/JobsCancelPost.json
 */
async function jobsCancelPost() {
  const subscriptionId = "fa68082f-8ff7-4a25-95c7-ce9da541242f";
  const resourceGroupName = "SdkRg5154";
  const jobName = "SdkJob952";
  const cancellationReason = { reason: "CancelTest" };
  const credential = new DefaultAzureCredential();
  const client = new DataBoxManagementClient(credential, subscriptionId);
  const result = await client.jobs.cancel(resourceGroupName, jobName, cancellationReason);
  console.log(result);
}

jobsCancelPost().catch(console.error);
