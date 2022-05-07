Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-resources-subscriptions_2.0.1/sdk/resources-subscriptions/arm-resources-subscriptions/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SubscriptionClient } = require("@azure/arm-resources-subscriptions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets details about a specified subscription.
 *
 * @summary Gets details about a specified subscription.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-01-01/examples/GetSubscription.json
 */
async function getASingleSubscription() {
  const subscriptionId = "291bba3f-e0a5-47bc-a099-3bdcb2a50a05";
  const credential = new DefaultAzureCredential();
  const client = new SubscriptionClient(credential);
  const result = await client.subscriptions.get(subscriptionId);
  console.log(result);
}

getASingleSubscription().catch(console.error);
```
