const { MicrosoftStorageSync } = require("@azure/arm-storagesync");
const { DefaultAzureCredential } = require("@azure/identity");

async function syncGroupsListByStorageSyncService() {
  const subscriptionId = "52b8da2f-61e0-4a1f-8dde-336911f367fb";
  const resourceGroupName = "SampleResourceGroup_1";
  const storageSyncServiceName = "SampleStorageSyncService_1";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftStorageSync(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.syncGroups.listByStorageSyncService(
    resourceGroupName,
    storageSyncServiceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

syncGroupsListByStorageSyncService().catch(console.error);
