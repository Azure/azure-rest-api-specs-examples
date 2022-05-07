Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-relay_3.0.1/sdk/relay/arm-relay/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { RelayAPI } = require("@azure/arm-relay");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the description for the specified namespace.
 *
 * @summary Returns the description for the specified namespace.
 * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/NameSpaces/RelayNameSpaceGet.json
 */
async function relayNameSpaceGet() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "resourcegroup";
  const namespaceName = "example-RelayNamespace-01";
  const credential = new DefaultAzureCredential();
  const client = new RelayAPI(credential, subscriptionId);
  const result = await client.namespaces.get(resourceGroupName, namespaceName);
  console.log(result);
}

relayNameSpaceGet().catch(console.error);
```
