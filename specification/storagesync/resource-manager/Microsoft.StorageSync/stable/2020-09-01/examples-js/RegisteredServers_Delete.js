const { MicrosoftStorageSync } = require("@azure/arm-storagesync");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete the given registered server.
 *
 * @summary Delete the given registered server.
 * x-ms-original-file: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/RegisteredServers_Delete.json
 */
async function registeredServersDelete() {
  const subscriptionId = "52b8da2f-61e0-4a1f-8dde-336911f367fb";
  const resourceGroupName = "SampleResourceGroup_1";
  const storageSyncServiceName = "SampleStorageSyncService_1";
  const serverId = "41166691-ab03-43e9-ab3e-0330eda162ac";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftStorageSync(credential, subscriptionId);
  const result = await client.registeredServers.beginDeleteAndWait(
    resourceGroupName,
    storageSyncServiceName,
    serverId
  );
  console.log(result);
}

registeredServersDelete().catch(console.error);
