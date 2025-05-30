const { DeviceRegistryManagementClient } = require("@azure/arm-deviceregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a Asset
 *
 * @summary update a Asset
 * x-ms-original-file: 2024-09-01-preview/Update_Asset.json
 */
async function updateAsset() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new DeviceRegistryManagementClient(credential, subscriptionId);
  const result = await client.assets.update("myResourceGroup", "my-asset", {
    properties: { enabled: false, displayName: "NewAssetDisplayName" },
  });
  console.log(result);
}
