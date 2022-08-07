const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an Asset Filter associated with the specified Asset.
 *
 * @summary Deletes an Asset Filter associated with the specified Asset.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/assetFilters-delete.json
 */
async function deleteAnAssetFilter() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contosomedia";
  const assetName = "ClimbingMountRainer";
  const filterName = "assetFilterWithTimeWindowAndTrack";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.assetFilters.delete(
    resourceGroupName,
    accountName,
    assetName,
    filterName
  );
  console.log(result);
}

deleteAnAssetFilter().catch(console.error);
