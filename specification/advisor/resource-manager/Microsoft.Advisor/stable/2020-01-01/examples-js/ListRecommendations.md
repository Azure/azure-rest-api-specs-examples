Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-advisor_3.0.1/sdk/advisor/arm-advisor/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { AdvisorManagementClient } = require("@azure/arm-advisor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Obtains cached recommendations for a subscription. The recommendations are generated or computed by invoking generateRecommendations.
 *
 * @summary Obtains cached recommendations for a subscription. The recommendations are generated or computed by invoking generateRecommendations.
 * x-ms-original-file: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/ListRecommendations.json
 */
async function listRecommendations() {
  const subscriptionId = "subscriptionId";
  const top = 10;
  const options = { top };
  const credential = new DefaultAzureCredential();
  const client = new AdvisorManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.recommendations.list(options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listRecommendations().catch(console.error);
```
