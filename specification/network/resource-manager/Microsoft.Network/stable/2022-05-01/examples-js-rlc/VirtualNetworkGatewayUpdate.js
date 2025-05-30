const createNetworkManagementClient = require("@azure-rest/arm-network").default,
  { getLongRunningPoller } = require("@azure-rest/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates a virtual network gateway in the specified resource group.
 *
 * @summary Creates or updates a virtual network gateway in the specified resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/VirtualNetworkGatewayUpdate.json
 */
async function updateVirtualNetworkGateway() {
  const credential = new DefaultAzureCredential();
  const client = createNetworkManagementClient(credential);
  const subscriptionId = "";
  const resourceGroupName = "rg1";
  const virtualNetworkGatewayName = "vpngw";
  const options = {
    body: {
      location: "centralus",
      properties: {
        activeActive: false,
        bgpSettings: {
          asn: 65515,
          bgpPeeringAddress: "10.0.1.30",
          peerWeight: 0,
        },
        customRoutes: { addressPrefixes: ["101.168.0.6/32"] },
        disableIPSecReplayProtection: false,
        enableBgp: false,
        enableBgpRouteTranslationForNat: false,
        enableDnsForwarding: true,
        gatewayType: "Vpn",
        ipConfigurations: [
          {
            name: "gwipconfig1",
            properties: {
              privateIPAllocationMethod: "Dynamic",
              publicIPAddress: {
                id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/gwpip",
              },
              subnet: {
                id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/GatewaySubnet",
              },
            },
          },
        ],
        natRules: [
          {
            name: "natRule1",
            id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/natRules/natRule1",
            properties: {
              type: "Static",
              externalMappings: [{ addressSpace: "50.0.0.0/24" }],
              internalMappings: [{ addressSpace: "10.10.0.0/24" }],
              ipConfigurationId: "",
              mode: "EgressSnat",
            },
          },
          {
            name: "natRule2",
            id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/natRules/natRule2",
            properties: {
              type: "Static",
              externalMappings: [{ addressSpace: "30.0.0.0/24" }],
              internalMappings: [{ addressSpace: "20.10.0.0/24" }],
              ipConfigurationId: "",
              mode: "IngressSnat",
            },
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
      },
    },
    queryParameters: { "api-version": "2022-05-01" },
  };
  const initialResponse = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways/{virtualNetworkGatewayName}",
      subscriptionId,
      resourceGroupName,
      virtualNetworkGatewayName,
    )
    .put(options);
  const poller = getLongRunningPoller(client, initialResponse);
  const result = await poller.pollUntilDone();
  console.log(result);
}

updateVirtualNetworkGateway().catch(console.error);
