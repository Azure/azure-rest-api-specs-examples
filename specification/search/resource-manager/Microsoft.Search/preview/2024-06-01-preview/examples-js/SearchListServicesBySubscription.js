const { SearchManagementClient } = require("@azure/arm-search");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of all Search services in the given subscription.
 *
 * @summary Gets a list of all Search services in the given subscription.
 * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/preview/2024-06-01-preview/examples/SearchListServicesBySubscription.json
 */
async function searchListServicesBySubscription() {
  const subscriptionId = process.env["SEARCH_SUBSCRIPTION_ID"] || "subid";
  const credential = new DefaultAzureCredential();
  const client = new SearchManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.services.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
