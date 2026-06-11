const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to delete a export.
 *
 * @summary the operation to delete a export.
 * x-ms-original-file: 2025-03-01/ExportDeleteByEnrollmentAccount.json
 */
async function exportDeleteByEnrollmentAccount() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  await client.exports.delete(
    "providers/Microsoft.Billing/billingAccounts/100/enrollmentAccounts/456",
    "TestExport",
  );
}
