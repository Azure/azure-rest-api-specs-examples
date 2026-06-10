const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to list all exports at the given scope.
 *
 * @summary the operation to list all exports at the given scope.
 * x-ms-original-file: 2025-03-01/ExportsGetByResourceGroup.json
 */
async function exportsGetByResourceGroup() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.exports.list(
    "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG",
  );
  console.log(result);
}
