const { SearchManagementClient } = require("@azure/arm-search");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the details of the private endpoint connection to the search service in the given resource group.
 *
 * @summary Gets the details of the private endpoint connection to the search service in the given resource group.
 * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/preview/2024-06-01-preview/examples/GetPrivateEndpointConnection.json
 */
async function privateEndpointConnectionGet() {
  const subscriptionId = process.env["SEARCH_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["SEARCH_RESOURCE_GROUP"] || "rg1";
  const searchServiceName = "mysearchservice";
  const privateEndpointConnectionName = "testEndpoint.50bf4fbe-d7c1-4b48-a642-4f5892642546";
  const credential = new DefaultAzureCredential();
  const client = new SearchManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.get(
    resourceGroupName,
    searchServiceName,
    privateEndpointConnectionName,
  );
  console.log(result);
}
