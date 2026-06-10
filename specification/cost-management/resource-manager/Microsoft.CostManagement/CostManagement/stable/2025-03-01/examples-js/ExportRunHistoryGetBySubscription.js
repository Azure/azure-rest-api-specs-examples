const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to get the run history of an export for the defined scope and export name.
 *
 * @summary the operation to get the run history of an export for the defined scope and export name.
 * x-ms-original-file: 2025-03-01/ExportRunHistoryGetBySubscription.json
 */
async function exportRunHistoryGetBySubscription() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.exports.getExecutionHistory(
    "subscriptions/00000000-0000-0000-0000-000000000000",
    "TestExport",
  );
  console.log(result);
}
