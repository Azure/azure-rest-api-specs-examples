const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the specified Virtual Router Peering.
 *
 * @summary Creates or updates the specified Virtual Router Peering.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VirtualRouterPeeringPut.json
 */
async function createVirtualRouterPeering() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualRouterName = "virtualRouter";
  const peeringName = "peering1";
  const parameters = {
    peerAsn: 20000,
    peerIp: "192.168.1.5",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualRouterPeerings.beginCreateOrUpdateAndWait(
    resourceGroupName,
    virtualRouterName,
    peeringName,
    parameters
  );
  console.log(result);
}

createVirtualRouterPeering().catch(console.error);
