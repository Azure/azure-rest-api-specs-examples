const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to run an export.
 *
 * @summary The operation to run an export.
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExportRunByBillingAccount.json
 */
async function exportRunByBillingAccount() {
  const scope = "providers/Microsoft.Billing/billingAccounts/123456";
  const exportName = "TestExport";
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.exports.execute(scope, exportName);
  console.log(result);
}
