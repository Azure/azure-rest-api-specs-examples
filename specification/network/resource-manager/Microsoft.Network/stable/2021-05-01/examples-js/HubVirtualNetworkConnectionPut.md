```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a hub virtual network connection if it doesn't exist else updates the existing one.
 *
 * @summary Creates a hub virtual network connection if it doesn't exist else updates the existing one.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/HubVirtualNetworkConnectionPut.json
 */
async function hubVirtualNetworkConnectionPut() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualHubName = "virtualHub1";
  const connectionName = "connection1";
  const hubVirtualNetworkConnectionParameters = {
    enableInternetSecurity: false,
    remoteVirtualNetwork: {
      id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/SpokeVnet1",
    },
    routingConfiguration: {
      associatedRouteTable: {
        id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1",
      },
      propagatedRouteTables: {
        ids: [
          {
            id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1",
          },
        ],
        labels: ["label1", "label2"],
      },
      vnetRoutes: {
        staticRoutes: [
          {
            name: "route1",
            addressPrefixes: ["10.1.0.0/16", "10.2.0.0/16"],
            nextHopIpAddress: "10.0.0.68",
          },
          {
            name: "route2",
            addressPrefixes: ["10.3.0.0/16", "10.4.0.0/16"],
            nextHopIpAddress: "10.0.0.65",
          },
        ],
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.hubVirtualNetworkConnections.beginCreateOrUpdateAndWait(
    resourceGroupName,
    virtualHubName,
    connectionName,
    hubVirtualNetworkConnectionParameters
  );
  console.log(result);
}

hubVirtualNetworkConnectionPut().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_27.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.
