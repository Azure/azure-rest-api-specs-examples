const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all authorizations in an express route port.
 *
 * @summary Gets all authorizations in an express route port.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/ExpressRoutePortAuthorizationList.json
 */
async function listExpressRoutePortAuthorization() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const expressRoutePortName = "expressRoutePortName";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.expressRoutePortAuthorizations.list(
    resourceGroupName,
    expressRoutePortName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listExpressRoutePortAuthorization().catch(console.error);
