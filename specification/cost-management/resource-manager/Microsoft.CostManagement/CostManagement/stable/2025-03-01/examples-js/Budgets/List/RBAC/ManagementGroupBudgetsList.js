const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists all budgets for the defined scope.
 *
 * @summary lists all budgets for the defined scope.
 * x-ms-original-file: 2025-03-01/Budgets/List/RBAC/ManagementGroupBudgetsList.json
 */
async function managementGroupBudgetsList() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const resArray = new Array();
  for await (const item of client.budgets.list(
    "Microsoft.Management/managementGroups/MYDEVTESTMG",
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
