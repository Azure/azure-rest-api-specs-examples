const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all authorizations in an express route port.
 *
 * @summary Gets all authorizations in an express route port.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/ExpressRoutePortAuthorizationList.json
 */
async function listExpressRoutePortAuthorization() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const expressRoutePortName = "expressRoutePortName";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.expressRoutePortAuthorizations.list(
    resourceGroupName,
    expressRoutePortName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
