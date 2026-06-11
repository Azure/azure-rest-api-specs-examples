const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to run an export.
 *
 * @summary the operation to run an export.
 * x-ms-original-file: 2025-03-01/ExportRunByDepartment.json
 */
async function exportRunByDepartment() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  await client.exports.execute(
    "providers/Microsoft.Billing/billingAccounts/12/departments/1234",
    "TestExport",
  );
}
