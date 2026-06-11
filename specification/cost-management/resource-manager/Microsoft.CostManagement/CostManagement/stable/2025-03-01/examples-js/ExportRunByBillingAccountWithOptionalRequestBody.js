const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to run an export.
 *
 * @summary the operation to run an export.
 * x-ms-original-file: 2025-03-01/ExportRunByBillingAccountWithOptionalRequestBody.json
 */
async function exportRunByBillingAccountWithOptionalRequestBody() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  await client.exports.execute("providers/Microsoft.Billing/billingAccounts/123456", "TestExport");
}
