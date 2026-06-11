const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to get the export for the defined scope by export name.
 *
 * @summary the operation to get the export for the defined scope by export name.
 * x-ms-original-file: 2025-03-01/ExportGetByEnrollmentAccount.json
 */
async function exportGetByEnrollmentAccount() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.exports.get(
    "providers/Microsoft.Billing/billingAccounts/100/enrollmentAccounts/456",
    "TestExport",
  );
  console.log(result);
}
