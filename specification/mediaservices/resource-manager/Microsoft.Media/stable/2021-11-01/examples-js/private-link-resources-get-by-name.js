const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get details of a group ID.
 *
 * @summary Get details of a group ID.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/private-link-resources-get-by-name.json
 */
async function getDetailsOfAGroupId() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contososports";
  const name = "keydelivery";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.privateLinkResources.get(resourceGroupName, accountName, name);
  console.log(result);
}

getDetailsOfAGroupId().catch(console.error);
