const { StorageImportExport } = require("@azure/arm-storageimportexport");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns all active and completed jobs in a subscription.
 *
 * @summary Returns all active and completed jobs in a subscription.
 * x-ms-original-file: specification/storageimportexport/resource-manager/Microsoft.ImportExport/preview/2021-01-01/examples/ListJobsInSubscription.json
 */
async function listJobsInASubscription() {
  const subscriptionId =
    process.env["STORAGEIMPORTEXPORT_SUBSCRIPTION_ID"] || "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx";
  const credential = new DefaultAzureCredential();
  const client = new StorageImportExport(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.jobs.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
