const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified authorization from the specified express route port.
 *
 * @summary Gets the specified authorization from the specified express route port.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/ExpressRoutePortAuthorizationGet.json
 */
async function getExpressRoutePortAuthorization() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const expressRoutePortName = "expressRoutePortName";
  const authorizationName = "authorizationName";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRoutePortAuthorizations.get(
    resourceGroupName,
    expressRoutePortName,
    authorizationName
  );
  console.log(result);
}

getExpressRoutePortAuthorization().catch(console.error);
