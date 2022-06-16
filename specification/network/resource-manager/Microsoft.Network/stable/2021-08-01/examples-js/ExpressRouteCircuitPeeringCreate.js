const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a peering in the specified express route circuits.
 *
 * @summary Creates or updates a peering in the specified express route circuits.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/ExpressRouteCircuitPeeringCreate.json
 */
async function createExpressRouteCircuitPeerings() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const circuitName = "circuitName";
  const peeringName = "AzurePrivatePeering";
  const peeringParameters = {
    peerASN: 200,
    primaryPeerAddressPrefix: "192.168.16.252/30",
    secondaryPeerAddressPrefix: "192.168.18.252/30",
    vlanId: 200,
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteCircuitPeerings.beginCreateOrUpdateAndWait(
    resourceGroupName,
    circuitName,
    peeringName,
    peeringParameters
  );
  console.log(result);
}

createExpressRouteCircuitPeerings().catch(console.error);
