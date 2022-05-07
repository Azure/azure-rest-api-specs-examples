Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_27.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a virtual wan p2s vpn gateway if it doesn't exist else updates the existing gateway.
 *
 * @summary Creates a virtual wan p2s vpn gateway if it doesn't exist else updates the existing gateway.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/P2SVpnGatewayPut.json
 */
async function p2SVpnGatewayPut() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const gatewayName = "p2sVpnGateway1";
  const p2SVpnGatewayParameters = {
    customDnsServers: ["1.1.1.1", "2.2.2.2"],
    isRoutingPreferenceInternet: false,
    location: "West US",
    p2SConnectionConfigurations: [
      {
        name: "P2SConnectionConfig1",
        id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/p2sVpnGateways/p2sVpnGateway1/p2sConnectionConfigurations/P2SConnectionConfig1",
        routingConfiguration: {
          associatedRouteTable: {
            id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1",
          },
          propagatedRouteTables: {
            ids: [
              {
                id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1",
              },
              {
                id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable2",
              },
              {
                id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable3",
              },
            ],
            labels: ["label1", "label2"],
          },
          vnetRoutes: { staticRoutes: [] },
        },
        vpnClientAddressPool: { addressPrefixes: ["101.3.0.0/16"] },
      },
    ],
    tags: { key1: "value1" },
    virtualHub: {
      id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1",
    },
    vpnGatewayScaleUnit: 1,
    vpnServerConfiguration: {
      id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnServerConfigurations/vpnServerConfiguration1",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.p2SVpnGateways.beginCreateOrUpdateAndWait(
    resourceGroupName,
    gatewayName,
    p2SVpnGatewayParameters
  );
  console.log(result);
}

p2SVpnGatewayPut().catch(console.error);
```
