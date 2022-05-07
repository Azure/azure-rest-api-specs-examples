Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-consumption_9.0.1/sdk/consumption/arm-consumption/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ConsumptionManagementClient } = require("@azure/arm-consumption");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the reservations summaries for the defined scope daily or monthly grain.
 *
 * @summary Lists the reservations summaries for the defined scope daily or monthly grain.
 * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ReservationSummariesMonthlyWithBillingProfileId.json
 */
async function reservationSummariesMonthlyWithBillingProfileId() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const scope = "providers/Microsoft.Billing/billingAccounts/12345:2468/billingProfiles/13579";
  const grain = "monthly";
  const credential = new DefaultAzureCredential();
  const client = new ConsumptionManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.reservationsSummaries.list(scope, grain)) {
    resArray.push(item);
  }
  console.log(resArray);
}

reservationSummariesMonthlyWithBillingProfileId().catch(console.error);
```
