```javascript
const { AdvisorManagementClient } = require("@azure/arm-advisor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Enables the snoozed or dismissed attribute of a recommendation. The snoozed or dismissed attribute is referred to as a suppression. Use this API to create or update the snoozed or dismissed status of a recommendation.
 *
 * @summary Enables the snoozed or dismissed attribute of a recommendation. The snoozed or dismissed attribute is referred to as a suppression. Use this API to create or update the snoozed or dismissed status of a recommendation.
 * x-ms-original-file: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/CreateSuppression.json
 */
async function createSuppression() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceUri = "resourceUri";
  const recommendationId = "recommendationId";
  const name = "suppressionName1";
  const suppressionContract = { ttl: "07:00:00:00" };
  const credential = new DefaultAzureCredential();
  const client = new AdvisorManagementClient(credential, subscriptionId);
  const result = await client.suppressions.create(
    resourceUri,
    recommendationId,
    name,
    suppressionContract
  );
  console.log(result);
}

createSuppression().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-advisor_3.0.1/sdk/advisor/arm-advisor/README.md) on how to add the SDK to your project and authenticate.
