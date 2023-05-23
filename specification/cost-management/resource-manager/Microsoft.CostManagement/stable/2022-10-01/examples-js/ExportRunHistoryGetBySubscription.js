const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to get the run history of an export for the defined scope and export name.
 *
 * @summary The operation to get the run history of an export for the defined scope and export name.
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExportRunHistoryGetBySubscription.json
 */
async function exportRunHistoryGetBySubscription() {
  const scope = "subscriptions/00000000-0000-0000-0000-000000000000";
  const exportName = "TestExport";
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.exports.getExecutionHistory(scope, exportName);
  console.log(result);
}
