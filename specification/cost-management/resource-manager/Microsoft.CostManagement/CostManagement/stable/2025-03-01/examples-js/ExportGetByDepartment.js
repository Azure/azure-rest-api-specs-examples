const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to get the export for the defined scope by export name.
 *
 * @summary the operation to get the export for the defined scope by export name.
 * x-ms-original-file: 2025-03-01/ExportGetByDepartment.json
 */
async function exportGetByDepartment() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.exports.get(
    "providers/Microsoft.Billing/billingAccounts/12/departments/1234",
    "TestExport",
  );
  console.log(result);
}
