const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all stats from an express route circuit in a resource group.
 *
 * @summary Gets all stats from an express route circuit in a resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ExpressRouteCircuitPeeringStats.json
 */
async function getExpressRouteCircuitPeeringTrafficStats() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const circuitName = "circuitName";
  const peeringName = "peeringName";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteCircuits.getPeeringStats(
    resourceGroupName,
    circuitName,
    peeringName
  );
  console.log(result);
}

getExpressRouteCircuitPeeringTrafficStats().catch(console.error);
