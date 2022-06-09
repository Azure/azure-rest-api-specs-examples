```javascript
const { FeatureClient } = require("@azure/arm-features");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the preview feature with the specified name.
 *
 * @summary Gets the preview feature with the specified name.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Features/stable/2021-07-01/examples/getFeature.json
 */
async function getFeature() {
  const subscriptionId = "subid";
  const resourceProviderNamespace = "Resource Provider Namespace";
  const featureName = "feature";
  const credential = new DefaultAzureCredential();
  const client = new FeatureClient(credential, subscriptionId);
  const result = await client.features.get(resourceProviderNamespace, featureName);
  console.log(result);
}

getFeature().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-features_3.0.1/sdk/features/arm-features/README.md) on how to add the SDK to your project and authenticate.
