const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a VirtualHubBgpConnection resource if it doesn't exist else updates the existing VirtualHubBgpConnection.
 *
 * @summary Creates a VirtualHubBgpConnection resource if it doesn't exist else updates the existing VirtualHubBgpConnection.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/VirtualHubBgpConnectionPut.json
 */
async function virtualHubRouteTableV2Put() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualHubName = "hub1";
  const connectionName = "conn1";
  const parameters = {
    hubVirtualNetworkConnection: {
      id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1/hubVirtualNetworkConnections/hubVnetConn1",
    },
    peerAsn: 20000,
    peerIp: "192.168.1.5",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualHubBgpConnection.beginCreateOrUpdateAndWait(
    resourceGroupName,
    virtualHubName,
    connectionName,
    parameters
  );
  console.log(result);
}

virtualHubRouteTableV2Put().catch(console.error);
