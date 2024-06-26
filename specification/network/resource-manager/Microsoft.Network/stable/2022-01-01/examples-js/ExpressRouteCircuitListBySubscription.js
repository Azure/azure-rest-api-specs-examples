const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the express route circuits in a subscription.
 *
 * @summary Gets all the express route circuits in a subscription.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/ExpressRouteCircuitListBySubscription.json
 */
async function listExpressRouteCircuitsInASubscription() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.expressRouteCircuits.listAll()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listExpressRouteCircuitsInASubscription().catch(console.error);
