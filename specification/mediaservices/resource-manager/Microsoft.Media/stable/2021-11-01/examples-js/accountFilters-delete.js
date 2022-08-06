const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an Account Filter in the Media Services account.
 *
 * @summary Deletes an Account Filter in the Media Services account.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/accountFilters-delete.json
 */
async function deleteAnAccountFilter() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contosomedia";
  const filterName = "accountFilterWithTimeWindowAndTrack";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.accountFilters.delete(resourceGroupName, accountName, filterName);
  console.log(result);
}

deleteAnAccountFilter().catch(console.error);
