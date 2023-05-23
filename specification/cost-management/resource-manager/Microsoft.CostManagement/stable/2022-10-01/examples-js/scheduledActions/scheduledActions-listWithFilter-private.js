const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all private scheduled actions.
 *
 * @summary List all private scheduled actions.
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/scheduledActions/scheduledActions-listWithFilter-private.json
 */
async function privateScheduledActionsListFilterByViewId() {
  const filter = "properties/viewId eq '/providers/Microsoft.CostManagement/views/swaggerExample'";
  const options = { filter };
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.scheduledActions.list(options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
