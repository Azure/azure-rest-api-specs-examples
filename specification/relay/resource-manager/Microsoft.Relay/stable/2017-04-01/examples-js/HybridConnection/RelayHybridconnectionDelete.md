Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-relay_3.0.1/sdk/relay/arm-relay/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { RelayAPI } = require("@azure/arm-relay");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a hybrid connection.
 *
 * @summary Deletes a hybrid connection.
 * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/HybridConnection/RelayHybridconnectionDelete.json
 */
async function relayHybridconnectionDelete() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "resourcegroup";
  const namespaceName = "example-RelayNamespace-01";
  const hybridConnectionName = "example-Relay-Hybrid-01";
  const credential = new DefaultAzureCredential();
  const client = new RelayAPI(credential, subscriptionId);
  const result = await client.hybridConnections.delete(
    resourceGroupName,
    namespaceName,
    hybridConnectionName
  );
  console.log(result);
}

relayHybridconnectionDelete().catch(console.error);
```
