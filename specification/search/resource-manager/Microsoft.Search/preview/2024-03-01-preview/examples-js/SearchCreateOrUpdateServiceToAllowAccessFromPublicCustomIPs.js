const { SearchManagementClient } = require("@azure/arm-search");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a search service in the given resource group. If the search service already exists, all properties will be updated with the given values.
 *
 * @summary Creates or updates a search service in the given resource group. If the search service already exists, all properties will be updated with the given values.
 * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/preview/2024-03-01-preview/examples/SearchCreateOrUpdateServiceToAllowAccessFromPublicCustomIPs.json
 */
async function searchCreateOrUpdateServiceToAllowAccessFromPublicCustomIPs() {
  const subscriptionId = process.env["SEARCH_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["SEARCH_RESOURCE_GROUP"] || "rg1";
  const searchServiceName = "mysearchservice";
  const service = {
    hostingMode: "default",
    location: "westus",
    networkRuleSet: {
      ipRules: [{ value: "123.4.5.6" }, { value: "123.4.6.0/18" }],
    },
    partitionCount: 1,
    replicaCount: 1,
    sku: { name: "standard" },
    tags: { appName: "My e-commerce app" },
  };
  const credential = new DefaultAzureCredential();
  const client = new SearchManagementClient(credential, subscriptionId);
  const result = await client.services.beginCreateOrUpdateAndWait(
    resourceGroupName,
    searchServiceName,
    service,
  );
  console.log(result);
}
