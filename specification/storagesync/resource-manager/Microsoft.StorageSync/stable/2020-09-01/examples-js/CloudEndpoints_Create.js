const { MicrosoftStorageSync } = require("@azure/arm-storagesync");
const { DefaultAzureCredential } = require("@azure/identity");

async function cloudEndpointsCreate() {
  const subscriptionId = "52b8da2f-61e0-4a1f-8dde-336911f367fb";
  const resourceGroupName = "SampleResourceGroup_1";
  const storageSyncServiceName = "SampleStorageSyncService_1";
  const syncGroupName = "SampleSyncGroup_1";
  const cloudEndpointName = "SampleCloudEndpoint_1";
  const parameters = {
    azureFileShareName: "cvcloud-afscv-0719-058-a94a1354-a1fd-4e9a-9a50-919fad8c4ba4",
    friendlyName: "ankushbsubscriptionmgmtmab",
    storageAccountResourceId:
      "/subscriptions/744f4d70-6d17-4921-8970-a765d14f763f/resourceGroups/tminienv59svc/providers/Microsoft.Storage/storageAccounts/tminienv59storage",
    storageAccountTenantId: '"72f988bf-86f1-41af-91ab-2d7cd011db47"',
  };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftStorageSync(credential, subscriptionId);
  const result = await client.cloudEndpoints.beginCreateAndWait(
    resourceGroupName,
    storageSyncServiceName,
    syncGroupName,
    cloudEndpointName,
    parameters
  );
  console.log(result);
}

cloudEndpointsCreate().catch(console.error);
