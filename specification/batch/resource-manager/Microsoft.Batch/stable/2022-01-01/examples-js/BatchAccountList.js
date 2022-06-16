const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about the Batch accounts associated with the subscription.
 *
 * @summary Gets information about the Batch accounts associated with the subscription.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/BatchAccountList.json
 */
async function batchAccountList() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.batchAccountOperations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

batchAccountList().catch(console.error);
