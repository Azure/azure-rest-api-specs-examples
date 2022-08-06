const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all authorizations in an express route circuit.
 *
 * @summary Gets all authorizations in an express route circuit.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/ExpressRouteCircuitAuthorizationList.json
 */
async function listExpressRouteCircuitAuthorization() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const circuitName = "circuitName";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.expressRouteCircuitAuthorizations.list(
    resourceGroupName,
    circuitName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listExpressRouteCircuitAuthorization().catch(console.error);
