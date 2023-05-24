const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the alerts for scope defined.
 *
 * @summary Lists the alerts for scope defined.
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/EnrollmentAccountAlerts.json
 */
async function enrollmentAccountAlerts() {
  const scope = "providers/Microsoft.Billing/billingAccounts/12345:6789/enrollmentAccounts/456";
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.alerts.list(scope);
  console.log(result);
}
