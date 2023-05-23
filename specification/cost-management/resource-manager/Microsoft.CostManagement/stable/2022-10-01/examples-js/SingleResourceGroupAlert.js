const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the alert for the scope by alert ID.
 *
 * @summary Gets the alert for the scope by alert ID.
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/SingleResourceGroupAlert.json
 */
async function singleResourceGroupAlerts() {
  const scope =
    "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/ScreenSharingTest-peer";
  const alertId = "22222222-2222-2222-2222-222222222222";
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.alerts.get(scope, alertId);
  console.log(result);
}
