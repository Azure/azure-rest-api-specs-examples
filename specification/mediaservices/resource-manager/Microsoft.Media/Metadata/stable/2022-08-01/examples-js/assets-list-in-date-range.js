const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List Assets in the Media Services account with optional filtering and ordering
 *
 * @summary List Assets in the Media Services account with optional filtering and ordering
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/assets-list-in-date-range.json
 */
async function listAssetCreatedInADateRange() {
  const subscriptionId =
    process.env["MEDIASERVICES_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["MEDIASERVICES_RESOURCE_GROUP"] || "contoso";
  const accountName = "contosomedia";
  const filter = "properties/created gt 2012-06-01 and properties/created lt 2013-07-01";
  const orderby = "properties/created";
  const options = { filter, orderby };
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.assets.list(resourceGroupName, accountName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
