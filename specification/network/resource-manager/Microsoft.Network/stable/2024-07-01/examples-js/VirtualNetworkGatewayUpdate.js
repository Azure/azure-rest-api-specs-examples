const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates a virtual network gateway in the specified resource group.
 *
 * @summary Creates or updates a virtual network gateway in the specified resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/VirtualNetworkGatewayUpdate.json
 */
async function updateVirtualNetworkGateway() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const virtualNetworkGatewayName = "vpngw";
  const parameters = {
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
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/gwpip",
        },
        subnet: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/GatewaySubnet",
        },
      },
    ],
    location: "centralus",
    natRules: [
      {
        name: "natRule1",
        typePropertiesType: "Static",
        externalMappings: [{ addressSpace: "50.0.0.0/24" }],
        id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/natRules/natRule1",
        internalMappings: [{ addressSpace: "10.10.0.0/24" }],
        ipConfigurationId: "",
        mode: "EgressSnat",
      },
      {
        name: "natRule2",
        typePropertiesType: "Static",
        externalMappings: [{ addressSpace: "30.0.0.0/24" }],
        id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/natRules/natRule2",
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
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkGateways.beginCreateOrUpdateAndWait(
    resourceGroupName,
    virtualNetworkGatewayName,
    parameters,
  );
  console.log(result);
}
