const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the Content Key Policies in the account
 *
 * @summary Lists the Content Key Policies in the account
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/content-key-policies-list-in-date-range.json
 */
async function listsContentKeyPoliciesWithCreatedAndLastModifiedFilters() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contosomedia";
  const filter = "properties/lastModified gt 2016-06-01 and properties/created lt 2013-07-01";
  const options = { filter };
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.contentKeyPolicies.list(resourceGroupName, accountName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listsContentKeyPoliciesWithCreatedAndLastModifiedFilters().catch(console.error);
