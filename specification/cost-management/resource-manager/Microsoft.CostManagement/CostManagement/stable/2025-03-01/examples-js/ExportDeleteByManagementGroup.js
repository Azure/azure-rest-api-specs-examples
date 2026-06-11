const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to delete a export.
 *
 * @summary the operation to delete a export.
 * x-ms-original-file: 2025-03-01/ExportDeleteByManagementGroup.json
 */
async function exportDeleteByManagementGroup() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  await client.exports.delete(
    "providers/Microsoft.Management/managementGroups/TestMG",
    "TestExport",
  );
}
