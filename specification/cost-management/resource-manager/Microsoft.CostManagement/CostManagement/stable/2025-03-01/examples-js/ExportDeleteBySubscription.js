const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to delete a export.
 *
 * @summary the operation to delete a export.
 * x-ms-original-file: 2025-03-01/ExportDeleteBySubscription.json
 */
async function exportDeleteBySubscription() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  await client.exports.delete("subscriptions/00000000-0000-0000-0000-000000000000", "TestExport");
}
