const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to get the export for the defined scope by export name.
 *
 * @summary the operation to get the export for the defined scope by export name.
 * x-ms-original-file: 2025-03-01/ExportGetByResourceGroup.json
 */
async function exportGetByResourceGroup() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.exports.get(
    "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG",
    "TestExport",
  );
  console.log(result);
}
