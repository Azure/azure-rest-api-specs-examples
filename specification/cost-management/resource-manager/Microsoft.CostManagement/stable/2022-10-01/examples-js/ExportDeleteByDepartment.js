const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to delete a export.
 *
 * @summary The operation to delete a export.
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExportDeleteByDepartment.json
 */
async function exportDeleteByDepartment() {
  const scope = "providers/Microsoft.Billing/billingAccounts/12/departments/1234";
  const exportName = "TestExport";
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.exports.delete(scope, exportName);
  console.log(result);
}
