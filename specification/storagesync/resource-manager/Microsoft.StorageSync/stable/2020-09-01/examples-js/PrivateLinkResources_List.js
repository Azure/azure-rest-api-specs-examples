const { MicrosoftStorageSync } = require("@azure/arm-storagesync");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the private link resources that need to be created for a storage sync service.
 *
 * @summary Gets the private link resources that need to be created for a storage sync service.
 * x-ms-original-file: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/PrivateLinkResources_List.json
 */
async function privateLinkResourcesList() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res6977";
  const storageSyncServiceName = "sss2527";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftStorageSync(credential, subscriptionId);
  const result = await client.privateLinkResources.listByStorageSyncService(
    resourceGroupName,
    storageSyncServiceName
  );
  console.log(result);
}

privateLinkResourcesList().catch(console.error);
