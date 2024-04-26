const { DeviceRegistryManagementClient } = require("@azure/arm-deviceregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a Asset
 *
 * @summary Get a Asset
 * x-ms-original-file: specification/deviceregistry/resource-manager/Microsoft.DeviceRegistry/preview/2023-11-01-preview/examples/Get_Asset.json
 */
async function getAnAsset() {
  const subscriptionId =
    process.env["DEVICEREGISTRY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["DEVICEREGISTRY_RESOURCE_GROUP"] || "myResourceGroup";
  const assetName = "my-asset";
  const credential = new DefaultAzureCredential();
  const client = new DeviceRegistryManagementClient(credential, subscriptionId);
  const result = await client.assets.get(resourceGroupName, assetName);
  console.log(result);
}
