const { MicrosoftStorageSync } = require("@azure/arm-storagesync");
const { DefaultAzureCredential } = require("@azure/identity");

async function serverEndpointsRecallAction() {
  const subscriptionId = "52b8da2f-61e0-4a1f-8dde-336911f367fb";
  const resourceGroupName = "SampleResourceGroup_1";
  const storageSyncServiceName = "SampleStorageSyncService_1";
  const syncGroupName = "SampleSyncGroup_1";
  const serverEndpointName = "SampleServerEndpoint_1";
  const parameters = { pattern: "", recallPath: "" };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftStorageSync(credential, subscriptionId);
  const result = await client.serverEndpoints.beginRecallActionAndWait(
    resourceGroupName,
    storageSyncServiceName,
    syncGroupName,
    serverEndpointName,
    parameters
  );
  console.log(result);
}

serverEndpointsRecallAction().catch(console.error);
