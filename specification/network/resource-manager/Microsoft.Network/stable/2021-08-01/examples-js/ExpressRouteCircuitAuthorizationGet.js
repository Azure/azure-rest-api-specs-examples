const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified authorization from the specified express route circuit.
 *
 * @summary Gets the specified authorization from the specified express route circuit.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/ExpressRouteCircuitAuthorizationGet.json
 */
async function getExpressRouteCircuitAuthorization() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const circuitName = "circuitName";
  const authorizationName = "authorizationName";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteCircuitAuthorizations.get(
    resourceGroupName,
    circuitName,
    authorizationName
  );
  console.log(result);
}

getExpressRouteCircuitAuthorization().catch(console.error);
