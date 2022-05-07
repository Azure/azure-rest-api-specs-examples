Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-advisor_3.0.1/sdk/advisor/arm-advisor/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { AdvisorManagementClient } = require("@azure/arm-advisor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the list of metadata entities.
 *
 * @summary Gets the list of metadata entities.
 * x-ms-original-file: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/ListRecommendationMetadata.json
 */
async function getMetadata() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new AdvisorManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.recommendationMetadata.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

getMetadata().catch(console.error);
```
