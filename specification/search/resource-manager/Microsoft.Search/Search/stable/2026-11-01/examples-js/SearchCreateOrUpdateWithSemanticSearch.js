const { SearchManagementClient } = require("@azure/arm-search");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a search service in the given resource group. If the search service already exists, all properties will be updated with the given values.
 *
 * @summary creates or updates a search service in the given resource group. If the search service already exists, all properties will be updated with the given values.
 * x-ms-original-file: 2026-11-01/SearchCreateOrUpdateWithSemanticSearch.json
 */
async function searchCreateOrUpdateWithSemanticSearch() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new SearchManagementClient(credential, subscriptionId);
  const result = await client.services.createOrUpdate("rg1", "mysearchservice", {
    location: "westus",
    computeType: "Default",
    hostingMode: "Default",
    partitionCount: 1,
    replicaCount: 3,
    semanticSearch: "free",
    sku: { name: "standard" },
    tags: { "app-name": "My e-commerce app" },
  });
  console.log(result);
}
