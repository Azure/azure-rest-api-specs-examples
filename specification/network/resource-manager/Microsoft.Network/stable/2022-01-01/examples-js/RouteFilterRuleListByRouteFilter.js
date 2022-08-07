const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all RouteFilterRules in a route filter.
 *
 * @summary Gets all RouteFilterRules in a route filter.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/RouteFilterRuleListByRouteFilter.json
 */
async function routeFilterRuleListByRouteFilter() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const routeFilterName = "filterName";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.routeFilterRules.listByRouteFilter(
    resourceGroupName,
    routeFilterName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

routeFilterRuleListByRouteFilter().catch(console.error);
