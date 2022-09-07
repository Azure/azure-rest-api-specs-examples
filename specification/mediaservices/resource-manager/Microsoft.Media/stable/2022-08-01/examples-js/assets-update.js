const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing Asset in the Media Services account
 *
 * @summary Updates an existing Asset in the Media Services account
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/assets-update.json
 */
async function updateAnAsset() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contosomedia";
  const assetName = "ClimbingMountBaker";
  const parameters = {
    description: "A documentary showing the ascent of Mount Baker in HD",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.assets.update(resourceGroupName, accountName, assetName, parameters);
  console.log(result);
}

updateAnAsset().catch(console.error);
