const { DeviceRegistryManagementClient } = require("@azure/arm-deviceregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get a Asset
 *
 * @summary get a Asset
 * x-ms-original-file: 2024-09-01-preview/Get_Asset_With_SyncStatus.json
 */
async function getAssetWithSyncStatus() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new DeviceRegistryManagementClient(credential, subscriptionId);
  const result = await client.assets.get("myResourceGroup", "my-asset");
  console.log(result);
}
