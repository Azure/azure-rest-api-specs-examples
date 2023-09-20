const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified peering for the express route circuit.
 *
 * @summary Gets the specified peering for the express route circuit.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/ExpressRouteCircuitPeeringGet.json
 */
async function getExpressRouteCircuitPeering() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const circuitName = "circuitName";
  const peeringName = "MicrosoftPeering";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteCircuitPeerings.get(
    resourceGroupName,
    circuitName,
    peeringName
  );
  console.log(result);
}
