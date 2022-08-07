const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a Media Services account
 *
 * @summary Deletes a Media Services account
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/accounts-delete.json
 */
async function deleteAMediaServicesAccount() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contososports";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.mediaservices.delete(resourceGroupName, accountName);
  console.log(result);
}

deleteAMediaServicesAccount().catch(console.error);
