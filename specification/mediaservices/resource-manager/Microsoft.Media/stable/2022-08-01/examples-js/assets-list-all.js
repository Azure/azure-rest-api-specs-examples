const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List Assets in the Media Services account with optional filtering and ordering
 *
 * @summary List Assets in the Media Services account with optional filtering and ordering
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/assets-list-all.json
 */
async function listAllAssets() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contosomedia";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.assets.list(resourceGroupName, accountName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllAssets().catch(console.error);
