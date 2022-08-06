const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists Streaming Locators which are associated with this asset.
 *
 * @summary Lists Streaming Locators which are associated with this asset.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/assets-list-streaming-locators.json
 */
async function listAssetSasUrLs() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contosomedia";
  const assetName = "ClimbingMountSaintHelens";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.assets.listStreamingLocators(
    resourceGroupName,
    accountName,
    assetName
  );
  console.log(result);
}

listAssetSasUrLs().catch(console.error);
