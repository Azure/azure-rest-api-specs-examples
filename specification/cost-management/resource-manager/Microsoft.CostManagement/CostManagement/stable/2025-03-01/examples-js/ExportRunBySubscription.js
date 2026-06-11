const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to run an export.
 *
 * @summary the operation to run an export.
 * x-ms-original-file: 2025-03-01/ExportRunBySubscription.json
 */
async function exportRunBySubscription() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  await client.exports.execute("subscriptions/00000000-0000-0000-0000-000000000000", "TestExport");
}
