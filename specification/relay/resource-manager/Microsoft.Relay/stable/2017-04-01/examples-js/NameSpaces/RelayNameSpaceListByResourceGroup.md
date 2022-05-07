Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-relay_3.0.1/sdk/relay/arm-relay/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { RelayAPI } = require("@azure/arm-relay");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the available namespaces within the ResourceGroup.
 *
 * @summary Lists all the available namespaces within the ResourceGroup.
 * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/NameSpaces/RelayNameSpaceListByResourceGroup.json
 */
async function relayNameSpaceListByResourceGroup() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = "resourcegroup";
  const credential = new DefaultAzureCredential();
  const client = new RelayAPI(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.namespaces.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

relayNameSpaceListByResourceGroup().catch(console.error);
```
