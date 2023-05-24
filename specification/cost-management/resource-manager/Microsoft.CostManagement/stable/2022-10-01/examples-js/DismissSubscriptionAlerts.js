const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Dismisses the specified alert
 *
 * @summary Dismisses the specified alert
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/DismissSubscriptionAlerts.json
 */
async function patchSubscriptionAlerts() {
  const scope = "subscriptions/00000000-0000-0000-0000-000000000000";
  const alertId = "22222222-2222-2222-2222-222222222222";
  const parameters = { status: "Dismissed" };
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.alerts.dismiss(scope, alertId, parameters);
  console.log(result);
}
