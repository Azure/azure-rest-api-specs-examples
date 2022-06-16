const { FeatureClient } = require("@azure/arm-features");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the preview features in a provider namespace that are available through AFEC for the subscription.
 *
 * @summary Gets all the preview features in a provider namespace that are available through AFEC for the subscription.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Features/stable/2021-07-01/examples/listProviderFeatures.json
 */
async function listProviderFeatures() {
  const subscriptionId = "subid";
  const resourceProviderNamespace = "Resource Provider Namespace";
  const credential = new DefaultAzureCredential();
  const client = new FeatureClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.features.list(resourceProviderNamespace)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listProviderFeatures().catch(console.error);
