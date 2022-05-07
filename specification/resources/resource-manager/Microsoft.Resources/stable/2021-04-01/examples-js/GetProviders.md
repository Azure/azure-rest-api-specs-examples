Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-resources_5.0.1/sdk/resources/arm-resources/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ResourceManagementClient } = require("@azure/arm-resources");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all resource providers for a subscription.
 *
 * @summary Gets all resource providers for a subscription.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/GetProviders.json
 */
async function getProviders() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new ResourceManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.providers.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

getProviders().catch(console.error);
```
