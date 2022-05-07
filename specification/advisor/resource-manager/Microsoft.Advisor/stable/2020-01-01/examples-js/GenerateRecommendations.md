Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-advisor_3.0.1/sdk/advisor/arm-advisor/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { AdvisorManagementClient } = require("@azure/arm-advisor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Initiates the recommendation generation or computation process for a subscription. This operation is asynchronous. The generated recommendations are stored in a cache in the Advisor service.
 *
 * @summary Initiates the recommendation generation or computation process for a subscription. This operation is asynchronous. The generated recommendations are stored in a cache in the Advisor service.
 * x-ms-original-file: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/GenerateRecommendations.json
 */
async function generateRecommendations() {
  const subscriptionId = "subscriptionId";
  const credential = new DefaultAzureCredential();
  const client = new AdvisorManagementClient(credential, subscriptionId);
  const result = await client.recommendations.generate();
  console.log(result);
}

generateRecommendations().catch(console.error);
```
