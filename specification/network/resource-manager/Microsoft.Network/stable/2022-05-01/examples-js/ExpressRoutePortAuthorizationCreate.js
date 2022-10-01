const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an authorization in the specified express route port.
 *
 * @summary Creates or updates an authorization in the specified express route port.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/ExpressRoutePortAuthorizationCreate.json
 */
async function createExpressRoutePortAuthorization() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const expressRoutePortName = "expressRoutePortName";
  const authorizationName = "authorizatinName";
  const authorizationParameters = {};
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRoutePortAuthorizations.beginCreateOrUpdateAndWait(
    resourceGroupName,
    expressRoutePortName,
    authorizationName,
    authorizationParameters
  );
  console.log(result);
}

createExpressRoutePortAuthorization().catch(console.error);
