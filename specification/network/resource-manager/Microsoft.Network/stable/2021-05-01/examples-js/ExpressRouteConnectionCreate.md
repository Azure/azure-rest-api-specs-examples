Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_27.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a connection between an ExpressRoute gateway and an ExpressRoute circuit.
 *
 * @summary Creates a connection between an ExpressRoute gateway and an ExpressRoute circuit.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ExpressRouteConnectionCreate.json
 */
async function expressRouteConnectionCreate() {
  const subscriptionId = "subid";
  const resourceGroupName = "resourceGroupName";
  const expressRouteGatewayName = "gateway-2";
  const connectionName = "connectionName";
  const putExpressRouteConnectionParameters = {
    name: "connectionName",
    authorizationKey: "authorizationKey",
    expressRouteCircuitPeering: {
      id: "/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteCircuits/circuitName/peerings/AzurePrivatePeering",
    },
    id: "/subscriptions/subid/resourceGroups/resourceGroupName/providers/Microsoft.Network/expressRouteGateways/gateway-2/expressRouteConnections/connectionName",
    routingWeight: 2,
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteConnections.beginCreateOrUpdateAndWait(
    resourceGroupName,
    expressRouteGatewayName,
    connectionName,
    putExpressRouteConnectionParameters
  );
  console.log(result);
}

expressRouteConnectionCreate().catch(console.error);
```
