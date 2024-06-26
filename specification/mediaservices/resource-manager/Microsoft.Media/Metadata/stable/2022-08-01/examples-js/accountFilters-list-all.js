const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List Account Filters in the Media Services account.
 *
 * @summary List Account Filters in the Media Services account.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/accountFilters-list-all.json
 */
async function listAllAccountFilters() {
  const subscriptionId =
    process.env["MEDIASERVICES_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["MEDIASERVICES_RESOURCE_GROUP"] || "contoso";
  const accountName = "contosomedia";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.accountFilters.list(resourceGroupName, accountName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
