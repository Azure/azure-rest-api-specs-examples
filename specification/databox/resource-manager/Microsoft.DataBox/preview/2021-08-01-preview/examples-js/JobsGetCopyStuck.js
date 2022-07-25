const { DataBoxManagementClient } = require("@azure/arm-databox");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about the specified job.
 *
 * @summary Gets information about the specified job.
 * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/JobsGetCopyStuck.json
 */
async function jobsGetCopyStuck() {
  const subscriptionId = "fa68082f-8ff7-4a25-95c7-ce9da541242f";
  const resourceGroupName = "dmstestresource";
  const jobName = "TJx-637505258985313014";
  const expand = "details";
  const options = { expand };
  const credential = new DefaultAzureCredential();
  const client = new DataBoxManagementClient(credential, subscriptionId);
  const result = await client.jobs.get(resourceGroupName, jobName, options);
  console.log(result);
}

jobsGetCopyStuck().catch(console.error);
