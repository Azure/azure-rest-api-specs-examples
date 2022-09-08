const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the details of an Account Filter in the Media Services account.
 *
 * @summary Get the details of an Account Filter in the Media Services account.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/accountFilters-get-by-name.json
 */
async function getAnAccountFilterByName() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contosomedia";
  const filterName = "accountFilterWithTrack";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.accountFilters.get(resourceGroupName, accountName, filterName);
  console.log(result);
}

getAnAccountFilterByName().catch(console.error);
