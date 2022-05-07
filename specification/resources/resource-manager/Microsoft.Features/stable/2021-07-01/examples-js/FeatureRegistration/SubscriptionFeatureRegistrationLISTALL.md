Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-features_3.0.1/sdk/features/arm-features/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { FeatureClient } = require("@azure/arm-features");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns subscription feature registrations for given subscription.
 *
 * @summary Returns subscription feature registrations for given subscription.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Features/stable/2021-07-01/examples/FeatureRegistration/SubscriptionFeatureRegistrationLISTALL.json
 */
async function getsAListOfFeatureRegistrations() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const credential = new DefaultAzureCredential();
  const client = new FeatureClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.subscriptionFeatureRegistrations.listAllBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

getsAListOfFeatureRegistrations().catch(console.error);
```
