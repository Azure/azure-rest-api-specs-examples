const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates tags of a route filter.
 *
 * @summary Updates tags of a route filter.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/RouteFilterUpdateTags.json
 */
async function updateRouteFilterTags() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const routeFilterName = "filterName";
  const parameters = { tags: { key1: "value1" } };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.routeFilters.updateTags(
    resourceGroupName,
    routeFilterName,
    parameters
  );
  console.log(result);
}

updateRouteFilterTags().catch(console.error);
