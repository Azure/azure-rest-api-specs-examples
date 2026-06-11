const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists the alerts for scope defined.
 *
 * @summary lists the alerts for scope defined.
 * x-ms-original-file: 2025-03-01/DepartmentAlerts.json
 */
async function departmentAlerts() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.alerts.list(
    "providers/Microsoft.Billing/billingAccounts/12345:6789/departments/123",
  );
  console.log(result);
}
