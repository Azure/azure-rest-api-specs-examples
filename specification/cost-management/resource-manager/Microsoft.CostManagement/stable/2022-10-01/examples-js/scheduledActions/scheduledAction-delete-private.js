const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a private scheduled action.
 *
 * @summary Delete a private scheduled action.
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/scheduledActions/scheduledAction-delete-private.json
 */
async function privateScheduledActionDelete() {
  const name = "monthlyCostByResource";
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.scheduledActions.delete(name);
  console.log(result);
}
