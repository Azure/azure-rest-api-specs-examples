const { MicrosoftStorageSync } = require("@azure/arm-storagesync");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a StorageSyncService list by subscription.
 *
 * @summary Get a StorageSyncService list by subscription.
 * x-ms-original-file: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/StorageSyncServices_ListBySubscription.json
 */
async function storageSyncServicesListBySubscription() {
  const subscriptionId = "52b8da2f-61e0-4a1f-8dde-336911f367fb";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftStorageSync(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.storageSyncServices.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

storageSyncServicesListBySubscription().catch(console.error);
