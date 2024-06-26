const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an Asset Filter associated with the specified Asset.
 *
 * @summary Deletes an Asset Filter associated with the specified Asset.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/assetFilters-delete.json
 */
async function deleteAnAssetFilter() {
  const subscriptionId =
    process.env["MEDIASERVICES_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["MEDIASERVICES_RESOURCE_GROUP"] || "contoso";
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
