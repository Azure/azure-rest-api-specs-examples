const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a Track in the asset
 *
 * @summary Deletes a Track in the asset
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/asset-tracks-delete.json
 */
async function deleteATrack() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contosomedia";
  const assetName = "ClimbingMountRainer";
  const trackName = "text2";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.tracks.beginDeleteAndWait(
    resourceGroupName,
    accountName,
    assetName,
    trackName
  );
  console.log(result);
}

deleteATrack().catch(console.error);
