const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists storage container URLs with shared access signatures (SAS) for uploading and downloading Asset content. The signatures are derived from the storage account keys.
 *
 * @summary Lists storage container URLs with shared access signatures (SAS) for uploading and downloading Asset content. The signatures are derived from the storage account keys.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/assets-list-sas-urls.json
 */
async function listAssetSasUrLs() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contosomedia";
  const assetName = "ClimbingMountBaker";
  const parameters = {
    expiryTime: new Date("2018-01-01T10:00:00.007Z"),
    permissions: "ReadWrite",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.assets.listContainerSas(
    resourceGroupName,
    accountName,
    assetName,
    parameters
  );
  console.log(result);
}

listAssetSasUrLs().catch(console.error);
