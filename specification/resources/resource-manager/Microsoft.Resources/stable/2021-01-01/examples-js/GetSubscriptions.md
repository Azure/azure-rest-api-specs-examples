Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-resources-subscriptions_2.0.1/sdk/resources-subscriptions/arm-resources-subscriptions/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SubscriptionClient } = require("@azure/arm-resources-subscriptions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all subscriptions for a tenant.
 *
 * @summary Gets all subscriptions for a tenant.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-01-01/examples/GetSubscriptions.json
 */
async function getAllSubscriptions() {
  const credential = new DefaultAzureCredential();
  const client = new SubscriptionClient(credential);
  const resArray = new Array();
  for await (let item of client.subscriptions.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

getAllSubscriptions().catch(console.error);
```
