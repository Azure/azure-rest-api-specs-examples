const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified route filter.
 *
 * @summary Gets the specified route filter.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/RouteFilterGet.json
 */
async function routeFilterGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const routeFilterName = "filterName";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.routeFilters.get(resourceGroupName, routeFilterName);
  console.log(result);
}

routeFilterGet().catch(console.error);
