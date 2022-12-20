const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of all the queues under the specified storage account
 *
 * @summary Gets a list of all the queues under the specified storage account
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/QueueOperationList.json
 */
async function queueOperationList() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res9290";
  const accountName = "sto328";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.queue.list(resourceGroupName, accountName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

queueOperationList().catch(console.error);
