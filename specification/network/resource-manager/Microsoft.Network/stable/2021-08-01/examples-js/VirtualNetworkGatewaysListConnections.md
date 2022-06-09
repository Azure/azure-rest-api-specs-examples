```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the connections in a virtual network gateway.
 *
 * @summary Gets all the connections in a virtual network gateway.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VirtualNetworkGatewaysListConnections.json
 */
async function virtualNetworkGatewaysListConnections() {
  const subscriptionId = "subid";
  const resourceGroupName = "testrg";
  const virtualNetworkGatewayName = "test-vpn-gateway-1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualNetworkGateways.listConnections(
    resourceGroupName,
    virtualNetworkGatewayName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

virtualNetworkGatewaysListConnections().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_28.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.
