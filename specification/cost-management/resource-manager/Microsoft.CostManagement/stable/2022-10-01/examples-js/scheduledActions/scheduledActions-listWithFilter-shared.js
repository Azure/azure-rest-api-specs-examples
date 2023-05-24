const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all shared scheduled actions within the given scope.
 *
 * @summary List all shared scheduled actions within the given scope.
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/scheduledActions/scheduledActions-listWithFilter-shared.json
 */
async function scheduledActionsListByScopeFilterByViewId() {
  const scope = "subscriptions/00000000-0000-0000-0000-000000000000";
  const filter = "properties/viewId eq '/providers/Microsoft.CostManagement/views/swaggerExample'";
  const options = { filter };
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.scheduledActions.listByScope(scope, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
