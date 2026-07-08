const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a virtual network gateway in the specified resource group.
 *
 * @summary creates or updates a virtual network gateway in the specified resource group.
 * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayUpdate.json
 */
async function updateVirtualNetworkGateway() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkGateways.createOrUpdate("rg1", "vpngw", {
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1":
          {},
      },
    },
    location: "centralus",
    active: false,
    allowRemoteVnetTraffic: false,
    allowVirtualWanTraffic: false,
    bgpSettings: { asn: 65515, bgpPeeringAddress: "10.0.1.30", peerWeight: 0 },
    customRoutes: { addressPrefixes: ["101.168.0.6/32"] },
    disableIPSecReplayProtection: false,
    enableBgp: false,
    enableBgpRouteTranslationForNat: false,
    enableDnsForwarding: true,
    enableHighBandwidthVpnGateway: false,
    gatewayType: "Vpn",
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
    natRules: [
      {
        id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/natRules/natRule1",
        typePropertiesType: "Static",
        externalMappings: [{ addressSpace: "50.0.0.0/24" }],
        internalMappings: [{ addressSpace: "10.10.0.0/24" }],
        ipConfigurationId: "",
        mode: "EgressSnat",
      },
      {
        id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/natRules/natRule2",
        typePropertiesType: "Static",
        externalMappings: [{ addressSpace: "30.0.0.0/24" }],
        internalMappings: [{ addressSpace: "20.10.0.0/24" }],
        ipConfigurationId: "",
        mode: "IngressSnat",
      },
    ],
    sku: { name: "VpnGw1", tier: "VpnGw1" },
    vpnClientConfiguration: {
      radiusServers: [
        {
          radiusServerAddress: "10.2.0.0",
          radiusServerScore: 20,
          radiusServerSecret: "radiusServerSecret",
        },
      ],
      vpnClientProtocols: ["OpenVPN"],
      vpnClientRevokedCertificates: [],
      vpnClientRootCertificates: [],
    },
    vpnType: "RouteBased",
  });
  console.log(result);
}
