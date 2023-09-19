const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all global reach peer connections associated with a private peering in an express route circuit.
 *
 * @summary Gets all global reach peer connections associated with a private peering in an express route circuit.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/PeerExpressRouteCircuitConnectionList.json
 */
async function listPeerExpressRouteCircuitConnection() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid1";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const circuitName = "ExpressRouteARMCircuitA";
  const peeringName = "AzurePrivatePeering";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.peerExpressRouteCircuitConnections.list(
    resourceGroupName,
    circuitName,
    peeringName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
