const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all route filters in a subscription.
 *
 * @summary Gets all route filters in a subscription.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/RouteFilterList.json
 */
async function routeFilterList() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.routeFilters.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

routeFilterList().catch(console.error);
