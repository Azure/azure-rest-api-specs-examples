const { DataBoxManagementClient } = require("@azure/arm-databox");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This method gets the unencrypted secrets related to the job.
 *
 * @summary This method gets the unencrypted secrets related to the job.
 * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/JobsListCredentials.json
 */
async function jobsListCredentials() {
  const subscriptionId = "fa68082f-8ff7-4a25-95c7-ce9da541242f";
  const resourceGroupName = "bvttoolrg6";
  const jobName = "TJ-636646322037905056";
  const credential = new DefaultAzureCredential();
  const client = new DataBoxManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.jobs.listCredentials(resourceGroupName, jobName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

jobsListCredentials().catch(console.error);
