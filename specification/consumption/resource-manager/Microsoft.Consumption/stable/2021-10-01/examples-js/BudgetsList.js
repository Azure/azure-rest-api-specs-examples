const { ConsumptionManagementClient } = require("@azure/arm-consumption");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all budgets for the defined scope.
 *
 * @summary Lists all budgets for the defined scope.
 * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/BudgetsList.json
 */
async function budgetsList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const scope = "subscriptions/00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new ConsumptionManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.budgets.list(scope)) {
    resArray.push(item);
  }
  console.log(resArray);
}

budgetsList().catch(console.error);
