const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the available express route service providers.
 *
 * @summary Gets all the available express route service providers.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/ExpressRouteProviderList.json
 */
async function listExpressRouteProviders() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.expressRouteServiceProviders.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listExpressRouteProviders().catch(console.error);
