const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve the ExpressRouteLink sub-resources of the specified ExpressRoutePort resource.
 *
 * @summary Retrieve the ExpressRouteLink sub-resources of the specified ExpressRoutePort resource.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/ExpressRouteLinkList.json
 */
async function expressRouteLinkGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const expressRoutePortName = "portName";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.expressRouteLinks.list(resourceGroupName, expressRoutePortName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

expressRouteLinkGet().catch(console.error);
