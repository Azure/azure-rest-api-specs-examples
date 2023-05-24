const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all private scheduled actions.
 *
 * @summary List all private scheduled actions.
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/scheduledActions/scheduledActions-list-private.json
 */
async function privateScheduledActionsList() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.scheduledActions.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
