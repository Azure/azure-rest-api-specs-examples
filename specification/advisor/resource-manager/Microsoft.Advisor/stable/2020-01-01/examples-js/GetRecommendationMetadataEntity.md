Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-advisor_3.0.1/sdk/advisor/arm-advisor/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { AdvisorManagementClient } = require("@azure/arm-advisor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the metadata entity.
 *
 * @summary Gets the metadata entity.
 * x-ms-original-file: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/GetRecommendationMetadataEntity.json
 */
async function getMetadata() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const name = "types";
  const credential = new DefaultAzureCredential();
  const client = new AdvisorManagementClient(credential, subscriptionId);
  const result = await client.recommendationMetadata.get(name);
  console.log(result);
}

getMetadata().catch(console.error);
```
