const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the Asset storage encryption keys used to decrypt content created by version 2 of the Media Services API
 *
 * @summary Gets the Asset storage encryption keys used to decrypt content created by version 2 of the Media Services API
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/assets-get-encryption-keys.json
 */
async function getAssetStorageEncryptionKeys() {
  const subscriptionId =
    process.env["MEDIASERVICES_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["MEDIASERVICES_RESOURCE_GROUP"] || "contoso";
  const accountName = "contosomedia";
  const assetName = "ClimbingMountSaintHelens";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.assets.getEncryptionKey(resourceGroupName, accountName, assetName);
  console.log(result);
}
