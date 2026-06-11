const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to get the run history of an export for the defined scope and export name.
 *
 * @summary the operation to get the run history of an export for the defined scope and export name.
 * x-ms-original-file: 2025-03-01/ExportRunHistoryGetByDepartment.json
 */
async function exportRunHistoryGetByDepartment() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.exports.getExecutionHistory(
    "providers/Microsoft.Billing/billingAccounts/12/departments/1234",
    "TestExport",
  );
  console.log(result);
}
