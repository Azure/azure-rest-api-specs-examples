const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified peering from the specified express route circuit.
 *
 * @summary Deletes the specified peering from the specified express route circuit.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2023-11-01/examples/ExpressRouteCircuitPeeringDelete.json
 */
async function deleteExpressRouteCircuitPeerings() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const circuitName = "circuitName";
  const peeringName = "peeringName";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteCircuitPeerings.beginDeleteAndWait(
    resourceGroupName,
    circuitName,
    peeringName,
  );
  console.log(result);
}
