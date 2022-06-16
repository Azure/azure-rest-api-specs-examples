const { SearchManagementClient } = require("@azure/arm-search");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a search service in the given resource group. If the search service already exists, all properties will be updated with the given values.
 *
 * @summary Creates or updates a search service in the given resource group. If the search service already exists, all properties will be updated with the given values.
 * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2020-08-01/examples/SearchCreateOrUpdateServiceWithIdentity.json
 */
async function searchCreateOrUpdateServiceWithIdentity() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const searchServiceName = "mysearchservice";
  const service = {
    hostingMode: "default",
    identity: { type: "SystemAssigned" },
    location: "westus",
    partitionCount: 1,
    replicaCount: 3,
    sku: { name: "standard" },
    tags: { appName: "My e-commerce app" },
  };
  const credential = new DefaultAzureCredential();
  const client = new SearchManagementClient(credential, subscriptionId);
  const result = await client.services.beginCreateOrUpdateAndWait(
    resourceGroupName,
    searchServiceName,
    service
  );
  console.log(result);
}

searchCreateOrUpdateServiceWithIdentity().catch(console.error);
