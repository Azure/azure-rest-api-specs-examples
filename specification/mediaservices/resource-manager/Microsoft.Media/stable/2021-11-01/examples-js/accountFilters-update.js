const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing Account Filter in the Media Services account.
 *
 * @summary Updates an existing Account Filter in the Media Services account.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/accountFilters-update.json
 */
async function updateAnAccountFilter() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contosomedia";
  const filterName = "accountFilterWithTimeWindowAndTrack";
  const parameters = {
    firstQuality: { bitrate: 128000 },
    presentationTimeRange: {
      endTimestamp: 170000000,
      forceEndTimestamp: false,
      liveBackoffDuration: 0,
      presentationWindowDuration: 9223372036854775000,
      startTimestamp: 10,
      timescale: 10000000,
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.accountFilters.update(
    resourceGroupName,
    accountName,
    filterName,
    parameters
  );
  console.log(result);
}

updateAnAccountFilter().catch(console.error);
