const { FeatureClient } = require("@azure/arm-features");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Registers the preview feature for the subscription.
 *
 * @summary Registers the preview feature for the subscription.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Features/stable/2021-07-01/examples/registerFeature.json
 */
async function registerFeature() {
  const subscriptionId = "subid";
  const resourceProviderNamespace = "Resource Provider Namespace";
  const featureName = "feature";
  const credential = new DefaultAzureCredential();
  const client = new FeatureClient(credential, subscriptionId);
  const result = await client.features.register(resourceProviderNamespace, featureName);
  console.log(result);
}

registerFeature().catch(console.error);
