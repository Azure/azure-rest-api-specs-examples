const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to list all exports at the given scope.
 *
 * @summary The operation to list all exports at the given scope.
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExportsGetByManagementGroup.json
 */
async function exportsGetByManagementGroup() {
  const scope = "providers/Microsoft.Management/managementGroups/TestMG";
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const result = await client.exports.list(scope);
  console.log(result);
}
