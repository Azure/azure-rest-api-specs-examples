const { SearchManagementClient } = require("@azure/arm-search");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Get a list of all Azure AI Search quota usages across the subscription.
 *
 * @summary Get a list of all Azure AI Search quota usages across the subscription.
 * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2025-05-01/examples/GetQuotaUsagesList.json
 */
async function getQuotaUsagesList() {
  const subscriptionId = process.env["SEARCH_SUBSCRIPTION_ID"] || "subid";
  const location = "westus";
  const credential = new DefaultAzureCredential();
  const client = new SearchManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.usages.listBySubscription(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}
