Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_27.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a vpn connection to a scalable vpn gateway if it doesn't exist else updates the existing connection.
 *
 * @summary Creates a vpn connection to a scalable vpn gateway if it doesn't exist else updates the existing connection.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VpnConnectionPut.json
 */
async function vpnConnectionPut() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const gatewayName = "gateway1";
  const connectionName = "vpnConnection1";
  const vpnConnectionParameters = {
    remoteVpnSite: {
      id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/vpnSite1",
    },
    trafficSelectorPolicies: [],
    vpnLinkConnections: [
      {
        name: "Connection-Link1",
        connectionBandwidth: 200,
        sharedKey: "key",
        usePolicyBasedTrafficSelectors: false,
        vpnConnectionProtocolType: "IKEv2",
        vpnLinkConnectionMode: "Default",
        vpnSiteLink: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/vpnSite1/vpnSiteLinks/siteLink1",
        },
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.vpnConnections.beginCreateOrUpdateAndWait(
    resourceGroupName,
    gatewayName,
    connectionName,
    vpnConnectionParameters
  );
  console.log(result);
}

vpnConnectionPut().catch(console.error);
```
