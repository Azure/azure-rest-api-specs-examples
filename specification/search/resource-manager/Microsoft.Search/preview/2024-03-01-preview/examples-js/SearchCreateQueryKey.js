const { SearchManagementClient } = require("@azure/arm-search");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Generates a new query key for the specified search service. You can create up to 50 query keys per service.
 *
 * @summary Generates a new query key for the specified search service. You can create up to 50 query keys per service.
 * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/preview/2024-03-01-preview/examples/SearchCreateQueryKey.json
 */
async function searchCreateQueryKey() {
  const subscriptionId = process.env["SEARCH_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["SEARCH_RESOURCE_GROUP"] || "rg1";
  const searchServiceName = "mysearchservice";
  const name = "An API key granting read-only access to the documents collection of an index.";
  const credential = new DefaultAzureCredential();
  const client = new SearchManagementClient(credential, subscriptionId);
  const result = await client.queryKeys.create(resourceGroupName, searchServiceName, name);
  console.log(result);
}
