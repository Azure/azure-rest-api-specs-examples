Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-consumption_9.0.1/sdk/consumption/arm-consumption/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ConsumptionManagementClient } = require("@azure/arm-consumption");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to delete a budget.
 *
 * @summary The operation to delete a budget.
 * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/DeleteBudget.json
 */
async function deleteBudget() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const scope = "subscriptions/00000000-0000-0000-0000-000000000000";
  const budgetName = "TestBudget";
  const credential = new DefaultAzureCredential();
  const client = new ConsumptionManagementClient(credential, subscriptionId);
  const result = await client.budgets.delete(scope, budgetName);
  console.log(result);
}

deleteBudget().catch(console.error);
```
