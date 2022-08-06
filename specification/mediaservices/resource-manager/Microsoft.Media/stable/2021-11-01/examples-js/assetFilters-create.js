const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an Asset Filter associated with the specified Asset.
 *
 * @summary Creates or updates an Asset Filter associated with the specified Asset.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/assetFilters-create.json
 */
async function createAnAssetFilter() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contosomedia";
  const assetName = "ClimbingMountRainer";
  const filterName = "newAssetFilter";
  const parameters = {
    firstQuality: { bitrate: 128000 },
    presentationTimeRange: {
      endTimestamp: 170000000,
      forceEndTimestamp: false,
      liveBackoffDuration: 0,
      presentationWindowDuration: 9223372036854775000,
      startTimestamp: 0,
      timescale: 10000000,
    },
    tracks: [
      {
        trackSelections: [
          { operation: "Equal", property: "Type", value: "Audio" },
          { operation: "NotEqual", property: "Language", value: "en" },
          { operation: "NotEqual", property: "FourCC", value: "EC-3" },
        ],
      },
      {
        trackSelections: [
          { operation: "Equal", property: "Type", value: "Video" },
          { operation: "Equal", property: "Bitrate", value: "3000000-5000000" },
        ],
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.assetFilters.createOrUpdate(
    resourceGroupName,
    accountName,
    assetName,
    filterName,
    parameters
  );
  console.log(result);
}

createAnAssetFilter().catch(console.error);
