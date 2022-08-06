const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a Job.
 *
 * @summary Deletes a Job.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/jobs-delete.json
 */
async function deleteAJob() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contosoresources";
  const accountName = "contosomedia";
  const transformName = "exampleTransform";
  const jobName = "jobToDelete";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.jobs.delete(resourceGroupName, accountName, transformName, jobName);
  console.log(result);
}

deleteAJob().catch(console.error);
