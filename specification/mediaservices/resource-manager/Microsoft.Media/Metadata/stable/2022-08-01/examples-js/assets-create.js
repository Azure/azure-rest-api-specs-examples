const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an Asset in the Media Services account
 *
 * @summary Creates or updates an Asset in the Media Services account
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/assets-create.json
 */
async function createAnAsset() {
  const subscriptionId =
    process.env["MEDIASERVICES_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["MEDIASERVICES_RESOURCE_GROUP"] || "contoso";
  const accountName = "contosomedia";
  const assetName = "ClimbingMountLogan";
  const parameters = {
    description: "A documentary showing the ascent of Mount Logan",
    storageAccountName: "storage0",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.assets.createOrUpdate(
    resourceGroupName,
    accountName,
    assetName,
    parameters
  );
  console.log(result);
}
