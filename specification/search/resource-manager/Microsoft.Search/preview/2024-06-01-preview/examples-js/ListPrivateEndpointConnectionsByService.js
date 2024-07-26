const { SearchManagementClient } = require("@azure/arm-search");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of all private endpoint connections in the given service.
 *
 * @summary Gets a list of all private endpoint connections in the given service.
 * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/preview/2024-06-01-preview/examples/ListPrivateEndpointConnectionsByService.json
 */
async function listPrivateEndpointConnectionsByService() {
  const subscriptionId = process.env["SEARCH_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["SEARCH_RESOURCE_GROUP"] || "rg1";
  const searchServiceName = "mysearchservice";
  const credential = new DefaultAzureCredential();
  const client = new SearchManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateEndpointConnections.listByService(
    resourceGroupName,
    searchServiceName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
