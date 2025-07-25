const { SearchManagementClient } = require("@azure/arm-search");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets the search service with the given name in the given resource group.
 *
 * @summary Gets the search service with the given name in the given resource group.
 * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2025-05-01/examples/SearchGetService.json
 */
async function searchGetService() {
  const subscriptionId = process.env["SEARCH_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["SEARCH_RESOURCE_GROUP"] || "rg1";
  const searchServiceName = "mysearchservice";
  const credential = new DefaultAzureCredential();
  const client = new SearchManagementClient(credential, subscriptionId);
  const result = await client.services.get(resourceGroupName, searchServiceName);
  console.log(result);
}
