const { SearchManagementClient } = require("@azure/arm-search");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the quota usage for a search sku in the given subscription.
 *
 * @summary Gets the quota usage for a search sku in the given subscription.
 * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/preview/2024-06-01-preview/examples/GetQuotaUsage.json
 */
async function getQuotaUsage() {
  const subscriptionId = process.env["SEARCH_SUBSCRIPTION_ID"] || "subid";
  const location = "westus";
  const skuName = "free";
  const credential = new DefaultAzureCredential();
  const client = new SearchManagementClient(credential, subscriptionId);
  const result = await client.usageBySubscriptionSku(location, skuName);
  console.log(result);
}
