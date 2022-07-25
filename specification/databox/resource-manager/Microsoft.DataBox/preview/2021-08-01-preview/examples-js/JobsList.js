const { DataBoxManagementClient } = require("@azure/arm-databox");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the jobs available under the subscription.
 *
 * @summary Lists all the jobs available under the subscription.
 * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/JobsList.json
 */
async function jobsList() {
  const subscriptionId = "fa68082f-8ff7-4a25-95c7-ce9da541242f";
  const credential = new DefaultAzureCredential();
  const client = new DataBoxManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.jobs.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

jobsList().catch(console.error);
