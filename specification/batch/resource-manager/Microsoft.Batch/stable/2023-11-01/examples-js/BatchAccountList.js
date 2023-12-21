const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about the Batch accounts associated with the subscription.
 *
 * @summary Gets information about the Batch accounts associated with the subscription.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-11-01/examples/BatchAccountList.json
 */
async function batchAccountList() {
  const subscriptionId = process.env["BATCH_SUBSCRIPTION_ID"] || "subid";
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.batchAccountOperations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
