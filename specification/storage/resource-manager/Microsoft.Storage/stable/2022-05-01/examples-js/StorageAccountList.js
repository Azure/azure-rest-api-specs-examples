const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the storage accounts available under the subscription. Note that storage keys are not returned; use the ListKeys operation for this.
 *
 * @summary Lists all the storage accounts available under the subscription. Note that storage keys are not returned; use the ListKeys operation for this.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/StorageAccountList.json
 */
async function storageAccountList() {
  const subscriptionId = "{subscription-id}";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.storageAccounts.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

storageAccountList().catch(console.error);
