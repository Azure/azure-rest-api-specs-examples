const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified rule from a route filter.
 *
 * @summary Gets the specified rule from a route filter.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/RouteFilterRuleGet.json
 */
async function routeFilterRuleGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const routeFilterName = "filterName";
  const ruleName = "filterName";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.routeFilterRules.get(resourceGroupName, routeFilterName, ruleName);
  console.log(result);
}

routeFilterRuleGet().catch(console.error);
