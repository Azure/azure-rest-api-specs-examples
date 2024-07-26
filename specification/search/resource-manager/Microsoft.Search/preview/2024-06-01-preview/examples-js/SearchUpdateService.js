const { SearchManagementClient } = require("@azure/arm-search");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing search service in the given resource group.
 *
 * @summary Updates an existing search service in the given resource group.
 * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/preview/2024-06-01-preview/examples/SearchUpdateService.json
 */
async function searchUpdateService() {
  const subscriptionId = process.env["SEARCH_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["SEARCH_RESOURCE_GROUP"] || "rg1";
  const searchServiceName = "mysearchservice";
  const service = {
    replicaCount: 2,
    tags: { appName: "My e-commerce app", newTag: "Adding a new tag" },
  };
  const credential = new DefaultAzureCredential();
  const client = new SearchManagementClient(credential, subscriptionId);
  const result = await client.services.update(resourceGroupName, searchServiceName, service);
  console.log(result);
}
