const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Track in the asset
 *
 * @summary Create or update a Track in the asset
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/asset-tracks-create.json
 */
async function createsATrack() {
  const subscriptionId =
    process.env["MEDIASERVICES_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["MEDIASERVICES_RESOURCE_GROUP"] || "contoso";
  const accountName = "contosomedia";
  const assetName = "ClimbingMountRainer";
  const trackName = "text3";
  const parameters = {
    track: {
      odataType: "#Microsoft.Media.TextTrack",
      displayName: "A new track",
      fileName: "text3.ttml",
      playerVisibility: "Visible",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.tracks.beginCreateOrUpdateAndWait(
    resourceGroupName,
    accountName,
    assetName,
    trackName,
    parameters
  );
  console.log(result);
}
