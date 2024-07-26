const { SearchManagementClient } = require("@azure/arm-search");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Regenerates either the primary or secondary admin API key. You can only regenerate one key at a time.
 *
 * @summary Regenerates either the primary or secondary admin API key. You can only regenerate one key at a time.
 * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/preview/2024-06-01-preview/examples/SearchRegenerateAdminKey.json
 */
async function searchRegenerateAdminKey() {
  const subscriptionId = process.env["SEARCH_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["SEARCH_RESOURCE_GROUP"] || "rg1";
  const searchServiceName = "mysearchservice";
  const keyKind = "primary";
  const credential = new DefaultAzureCredential();
  const client = new SearchManagementClient(credential, subscriptionId);
  const result = await client.adminKeys.regenerate(resourceGroupName, searchServiceName, keyKind);
  console.log(result);
}
