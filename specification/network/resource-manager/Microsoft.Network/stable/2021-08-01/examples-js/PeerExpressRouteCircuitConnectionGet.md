Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_28.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified Peer Express Route Circuit Connection from the specified express route circuit.
 *
 * @summary Gets the specified Peer Express Route Circuit Connection from the specified express route circuit.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/PeerExpressRouteCircuitConnectionGet.json
 */
async function peerExpressRouteCircuitConnectionGet() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const circuitName = "ExpressRouteARMCircuitA";
  const peeringName = "AzurePrivatePeering";
  const connectionName = "60aee347-e889-4a42-8c1b-0aae8b1e4013";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.peerExpressRouteCircuitConnections.get(
    resourceGroupName,
    circuitName,
    peeringName,
    connectionName
  );
  console.log(result);
}

peerExpressRouteCircuitConnectionGet().catch(console.error);
```
