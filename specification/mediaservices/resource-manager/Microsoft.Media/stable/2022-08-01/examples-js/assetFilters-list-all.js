const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List Asset Filters associated with the specified Asset.
 *
 * @summary List Asset Filters associated with the specified Asset.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/assetFilters-list-all.json
 */
async function listAllAssetFilters() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contosomedia";
  const assetName = "ClimbingMountRainer";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.assetFilters.list(resourceGroupName, accountName, assetName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllAssetFilters().catch(console.error);
