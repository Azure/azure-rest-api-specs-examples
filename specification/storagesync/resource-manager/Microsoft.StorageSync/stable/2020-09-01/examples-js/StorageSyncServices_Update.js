const { MicrosoftStorageSync } = require("@azure/arm-storagesync");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patch a given StorageSyncService.
 *
 * @summary Patch a given StorageSyncService.
 * x-ms-original-file: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/StorageSyncServices_Update.json
 */
async function storageSyncServicesUpdate() {
  const subscriptionId = "52b8da2f-61e0-4a1f-8dde-336911f367fb";
  const resourceGroupName = "SampleResourceGroup_1";
  const storageSyncServiceName = "SampleStorageSyncService_1";
  const parameters = {
    incomingTrafficPolicy: "AllowAllTraffic",
    tags: { dept: "IT", environment: "Test" },
  };
  const options = { parameters };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftStorageSync(credential, subscriptionId);
  const result = await client.storageSyncServices.beginUpdateAndWait(
    resourceGroupName,
    storageSyncServiceName,
    options
  );
  console.log(result);
}

storageSyncServicesUpdate().catch(console.error);
