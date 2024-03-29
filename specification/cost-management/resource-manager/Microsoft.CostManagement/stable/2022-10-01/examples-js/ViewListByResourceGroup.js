const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all views at the given scope.
 *
 * @summary Lists all views at the given scope.
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ViewListByResourceGroup.json
 */
async function resourceGroupViewList() {
  const scope = "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG";
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.views.listByScope(scope)) {
    resArray.push(item);
  }
  console.log(resArray);
}
