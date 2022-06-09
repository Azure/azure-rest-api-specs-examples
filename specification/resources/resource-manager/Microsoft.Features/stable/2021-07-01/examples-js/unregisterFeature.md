```javascript
const { FeatureClient } = require("@azure/arm-features");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Unregisters the preview feature for the subscription.
 *
 * @summary Unregisters the preview feature for the subscription.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Features/stable/2021-07-01/examples/unregisterFeature.json
 */
async function registerFeature() {
  const subscriptionId = "ff23096b-f5a2-46ea-bd62-59c3e93fef9a";
  const resourceProviderNamespace = "Resource Provider Namespace";
  const featureName = "feature";
  const credential = new DefaultAzureCredential();
  const client = new FeatureClient(credential, subscriptionId);
  const result = await client.features.unregister(resourceProviderNamespace, featureName);
  console.log(result);
}

registerFeature().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-features_3.0.1/sdk/features/arm-features/README.md) on how to add the SDK to your project and authenticate.
