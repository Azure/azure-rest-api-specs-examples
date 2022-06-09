```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an authorization in the specified express route circuit.
 *
 * @summary Creates or updates an authorization in the specified express route circuit.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ExpressRouteCircuitAuthorizationCreate.json
 */
async function createExpressRouteCircuitAuthorization() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const circuitName = "circuitName";
  const authorizationName = "authorizatinName";
  const authorizationParameters = {};
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.expressRouteCircuitAuthorizations.beginCreateOrUpdateAndWait(
    resourceGroupName,
    circuitName,
    authorizationName,
    authorizationParameters
  );
  console.log(result);
}

createExpressRouteCircuitAuthorization().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_27.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.
