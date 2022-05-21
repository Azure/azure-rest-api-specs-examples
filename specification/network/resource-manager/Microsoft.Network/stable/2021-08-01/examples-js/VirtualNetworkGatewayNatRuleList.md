Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_28.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves all nat rules for a particular virtual network gateway.
 *
 * @summary Retrieves all nat rules for a particular virtual network gateway.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VirtualNetworkGatewayNatRuleList.json
 */
async function virtualNetworkGatewayNatRuleList() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualNetworkGatewayName = "gateway1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualNetworkGatewayNatRules.listByVirtualNetworkGateway(
    resourceGroupName,
    virtualNetworkGatewayName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

virtualNetworkGatewayNatRuleList().catch(console.error);
```
