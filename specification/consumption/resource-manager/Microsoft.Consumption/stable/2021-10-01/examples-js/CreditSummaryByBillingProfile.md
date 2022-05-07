Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-consumption_9.0.1/sdk/consumption/arm-consumption/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ConsumptionManagementClient } = require("@azure/arm-consumption");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The credit summary by billingAccountId and billingProfileId.
 *
 * @summary The credit summary by billingAccountId and billingProfileId.
 * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/CreditSummaryByBillingProfile.json
 */
async function creditSummaryByBillingProfile() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const billingAccountId = "1234:5678";
  const billingProfileId = "2468";
  const credential = new DefaultAzureCredential();
  const client = new ConsumptionManagementClient(credential, subscriptionId);
  const result = await client.credits.get(billingAccountId, billingProfileId);
  console.log(result);
}

creditSummaryByBillingProfile().catch(console.error);
```
