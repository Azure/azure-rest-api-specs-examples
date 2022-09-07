const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an Account Filter in the Media Services account.
 *
 * @summary Creates or updates an Account Filter in the Media Services account.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/accountFilters-create.json
 */
async function createAnAccountFilter() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contosomedia";
  const filterName = "newAccountFilter";
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
  const result = await client.accountFilters.createOrUpdate(
    resourceGroupName,
    accountName,
    filterName,
    parameters
  );
  console.log(result);
}

createAnAccountFilter().catch(console.error);
