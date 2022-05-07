Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-consumption_9.0.1/sdk/consumption/arm-consumption/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ConsumptionManagementClient } = require("@azure/arm-consumption");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the balances for a scope by billing period and billingAccountId. Balances are available via this API only for May 1, 2014 or later.
 *
 * @summary Gets the balances for a scope by billing period and billingAccountId. Balances are available via this API only for May 1, 2014 or later.
 * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/BalancesByBillingAccountForBillingPeriod.json
 */
async function balances() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const billingAccountId = "123456";
  const billingPeriodName = "201702";
  const credential = new DefaultAzureCredential();
  const client = new ConsumptionManagementClient(credential, subscriptionId);
  const result = await client.balances.getForBillingPeriodByBillingAccount(
    billingAccountId,
    billingPeriodName
  );
  console.log(result);
}

balances().catch(console.error);
```
