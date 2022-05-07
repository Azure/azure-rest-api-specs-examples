Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-consumption_9.0.1/sdk/consumption/arm-consumption/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ConsumptionManagementClient } = require("@azure/arm-consumption");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Provides the aggregate cost of a management group and all child management groups by current billing period.
 *
 * @summary Provides the aggregate cost of a management group and all child management groups by current billing period.
 * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/AggregatedCostByManagementGroupFilterByDate.json
 */
async function aggregatedCostByManagementGroupFilterByDate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const managementGroupId = "managementGroupForTest";
  const filter = "usageStart ge '2018-08-15' and properties/usageStart le '2018-08-31'";
  const options = { filter };
  const credential = new DefaultAzureCredential();
  const client = new ConsumptionManagementClient(credential, subscriptionId);
  const result = await client.aggregatedCost.getByManagementGroup(managementGroupId, options);
  console.log(result);
}

aggregatedCostByManagementGroupFilterByDate().catch(console.error);
```
