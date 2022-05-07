Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-consumption_9.0.1/sdk/consumption/arm-consumption/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ConsumptionManagementClient } = require("@azure/arm-consumption");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Provides the aggregate cost of a management group and all child management groups by specified billing period
 *
 * @summary Provides the aggregate cost of a management group and all child management groups by specified billing period
 * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/AggregatedCostForBillingPeriodByManagementGroup.json
 */
async function aggregatedCostListForBillingPeriodByManagementGroup() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const managementGroupId = "managementGroupForTest";
  const billingPeriodName = "201807";
  const credential = new DefaultAzureCredential();
  const client = new ConsumptionManagementClient(credential, subscriptionId);
  const result = await client.aggregatedCost.getForBillingPeriodByManagementGroup(
    managementGroupId,
    billingPeriodName
  );
  console.log(result);
}

aggregatedCostListForBillingPeriodByManagementGroup().catch(console.error);
```
