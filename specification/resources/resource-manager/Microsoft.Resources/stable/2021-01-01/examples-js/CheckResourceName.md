```javascript
const { SubscriptionClient } = require("@azure/arm-resources-subscriptions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to A resource name is valid if it is not a reserved word, does not contains a reserved word and does not start with a reserved word
 *
 * @summary A resource name is valid if it is not a reserved word, does not contains a reserved word and does not start with a reserved word
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-01-01/examples/CheckResourceName.json
 */
async function checkValidityForAResourceName() {
  const credential = new DefaultAzureCredential();
  const client = new SubscriptionClient(credential);
  const result = await client.checkResourceName();
  console.log(result);
}

checkValidityForAResourceName().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-resources-subscriptions_2.0.1/sdk/resources-subscriptions/arm-resources-subscriptions/README.md) on how to add the SDK to your project and authenticate.
