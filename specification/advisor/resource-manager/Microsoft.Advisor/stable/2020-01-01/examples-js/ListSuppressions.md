Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-advisor_3.0.1/sdk/advisor/arm-advisor/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { AdvisorManagementClient } = require("@azure/arm-advisor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves the list of snoozed or dismissed suppressions for a subscription. The snoozed or dismissed attribute of a recommendation is referred to as a suppression.
 *
 * @summary Retrieves the list of snoozed or dismissed suppressions for a subscription. The snoozed or dismissed attribute of a recommendation is referred to as a suppression.
 * x-ms-original-file: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/ListSuppressions.json
 */
async function listSuppressions() {
  const subscriptionId = "subscriptionId1";
  const credential = new DefaultAzureCredential();
  const client = new AdvisorManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.suppressions.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listSuppressions().catch(console.error);
```
