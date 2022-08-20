const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists deleted accounts under the subscription.
 *
 * @summary Lists deleted accounts under the subscription.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/DeletedAccountList.json
 */
async function deletedAccountList() {
  const subscriptionId = "{subscription-id}";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.deletedAccounts.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

deletedAccountList().catch(console.error);
