const { MicrosoftStorageSync } = require("@azure/arm-storagesync");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patch a given ServerEndpoint.
 *
 * @summary Patch a given ServerEndpoint.
 * x-ms-original-file: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/ServerEndpoints_Update.json
 */
async function serverEndpointsUpdate() {
  const subscriptionId = "52b8da2f-61e0-4a1f-8dde-336911f367fb";
  const resourceGroupName = "SampleResourceGroup_1";
  const storageSyncServiceName = "SampleStorageSyncService_1";
  const syncGroupName = "SampleSyncGroup_1";
  const serverEndpointName = "SampleServerEndpoint_1";
  const parameters = {
    cloudTiering: "off",
    localCacheMode: "UpdateLocallyCachedFiles",
    offlineDataTransfer: "off",
    tierFilesOlderThanDays: 0,
    volumeFreeSpacePercent: 100,
  };
  const options = { parameters };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftStorageSync(credential, subscriptionId);
  const result = await client.serverEndpoints.beginUpdateAndWait(
    resourceGroupName,
    storageSyncServiceName,
    syncGroupName,
    serverEndpointName,
    options
  );
  console.log(result);
}

serverEndpointsUpdate().catch(console.error);
