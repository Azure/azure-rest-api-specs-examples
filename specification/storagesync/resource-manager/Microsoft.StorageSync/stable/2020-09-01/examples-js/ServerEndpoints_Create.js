const { MicrosoftStorageSync } = require("@azure/arm-storagesync");
const { DefaultAzureCredential } = require("@azure/identity");

async function serverEndpointsCreate() {
  const subscriptionId = "52b8da2f-61e0-4a1f-8dde-336911f367fb";
  const resourceGroupName = "SampleResourceGroup_1";
  const storageSyncServiceName = "SampleStorageSyncService_1";
  const syncGroupName = "SampleSyncGroup_1";
  const serverEndpointName = "SampleServerEndpoint_1";
  const parameters = {
    cloudTiering: "off",
    initialDownloadPolicy: "NamespaceThenModifiedFiles",
    initialUploadPolicy: "ServerAuthoritative",
    localCacheMode: "UpdateLocallyCachedFiles",
    offlineDataTransfer: "on",
    offlineDataTransferShareName: "myfileshare",
    serverLocalPath: "D:SampleServerEndpoint_1",
    serverResourceId:
      "/subscriptions/52b8da2f-61e0-4a1f-8dde-336911f367fb/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1/registeredServers/080d4133-bdb5-40a0-96a0-71a6057bfe9a",
    tierFilesOlderThanDays: 0,
    volumeFreeSpacePercent: 100,
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftStorageSync(credential, subscriptionId);
  const result = await client.serverEndpoints.beginCreateAndWait(
    resourceGroupName,
    storageSyncServiceName,
    syncGroupName,
    serverEndpointName,
    parameters
  );
  console.log(result);
}

serverEndpointsCreate().catch(console.error);
