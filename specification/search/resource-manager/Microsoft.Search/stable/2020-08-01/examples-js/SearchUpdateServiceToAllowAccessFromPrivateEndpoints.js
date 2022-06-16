const { SearchManagementClient } = require("@azure/arm-search");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing search service in the given resource group.
 *
 * @summary Updates an existing search service in the given resource group.
 * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2020-08-01/examples/SearchUpdateServiceToAllowAccessFromPrivateEndpoints.json
 */
async function searchUpdateServiceToAllowAccessFromPrivateEndpoints() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const searchServiceName = "mysearchservice";
  const service = {
    partitionCount: 1,
    publicNetworkAccess: "disabled",
    replicaCount: 1,
  };
  const credential = new DefaultAzureCredential();
  const client = new SearchManagementClient(credential, subscriptionId);
  const result = await client.services.update(resourceGroupName, searchServiceName, service);
  console.log(result);
}

searchUpdateServiceToAllowAccessFromPrivateEndpoints().catch(console.error);
