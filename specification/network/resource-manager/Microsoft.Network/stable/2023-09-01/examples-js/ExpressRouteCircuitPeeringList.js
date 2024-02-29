const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all peerings in a specified express route circuit.
 *
 * @summary Gets all peerings in a specified express route circuit.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-09-01/examples/ExpressRouteCircuitPeeringList.json
 */
async function listExpressRouteCircuitPeerings() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const circuitName = "circuitName";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.expressRouteCircuitPeerings.list(resourceGroupName, circuitName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
