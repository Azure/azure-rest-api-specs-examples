const { SearchManagementClient } = require("@azure/arm-search");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of all search services in the given subscription.
 *
 * @summary Gets a list of all search services in the given subscription.
 * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2020-08-01/examples/SearchListServicesBySubscription.json
 */
async function searchListServicesBySubscription() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new SearchManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.services.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

searchListServicesBySubscription().catch(console.error);
