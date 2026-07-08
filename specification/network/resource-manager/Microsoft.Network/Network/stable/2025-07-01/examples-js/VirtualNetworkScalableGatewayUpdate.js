const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a virtual network gateway in the specified resource group.
 *
 * @summary creates or updates a virtual network gateway in the specified resource group.
 * x-ms-original-file: 2025-07-01/VirtualNetworkScalableGatewayUpdate.json
 */
async function updateVirtualNetworkScalableGateway() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkGateways.createOrUpdate("rg1", "ergw", {
    location: "centralus",
    active: false,
    adminState: "Enabled",
    allowRemoteVnetTraffic: false,
    allowVirtualWanTraffic: false,
    autoScaleConfiguration: { bounds: { max: 3, min: 2 } },
    disableIPSecReplayProtection: false,
    enableBgp: false,
    enableBgpRouteTranslationForNat: false,
    gatewayType: "ExpressRoute",
    ipConfigurations: [
      {
        name: "gwipconfig1",
        privateIPAllocationMethod: "Dynamic",
        publicIPAddress: {
          id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/gwpip",
        },
        subnet: {
          id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/GatewaySubnet",
        },
      },
    ],
    natRules: [],
    sku: { name: "ErGwScale", tier: "ErGwScale" },
    virtualNetworkGatewayPolicyGroups: [],
    vpnType: "PolicyBased",
  });
  console.log(result);
}
