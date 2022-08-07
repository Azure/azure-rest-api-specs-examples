const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the details of an Asset Filter associated with the specified Asset.
 *
 * @summary Get the details of an Asset Filter associated with the specified Asset.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/assetFilters-get-by-name.json
 */
async function getAnAssetFilterByName() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contosomedia";
  const assetName = "ClimbingMountRainer";
  const filterName = "assetFilterWithTimeWindowAndTrack";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.assetFilters.get(
    resourceGroupName,
    accountName,
    assetName,
    filterName
  );
  console.log(result);
}

getAnAssetFilterByName().catch(console.error);
