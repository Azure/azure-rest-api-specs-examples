Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-features_3.0.1/sdk/features/arm-features/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { FeatureClient } = require("@azure/arm-features");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a feature registration.
 *
 * @summary Create or update a feature registration.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Features/stable/2021-07-01/examples/FeatureRegistration/SubscriptionFeatureRegistrationPUT.json
 */
async function createsAFeatureRegistration() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const providerNamespace = "subscriptionFeatureRegistrationGroupTestRG";
  const featureName = "testFeature";
  const subscriptionFeatureRegistrationType = {
    properties: {},
  };
  const options = {
    subscriptionFeatureRegistrationType,
  };
  const credential = new DefaultAzureCredential();
  const client = new FeatureClient(credential, subscriptionId);
  const result = await client.subscriptionFeatureRegistrations.createOrUpdate(
    providerNamespace,
    featureName,
    options
  );
  console.log(result);
}

createsAFeatureRegistration().catch(console.error);
```
