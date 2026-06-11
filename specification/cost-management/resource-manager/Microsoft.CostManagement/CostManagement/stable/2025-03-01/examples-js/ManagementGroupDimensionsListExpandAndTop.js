const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists the dimensions by the defined scope.
 *
 * @summary lists the dimensions by the defined scope.
 * x-ms-original-file: 2025-03-01/ManagementGroupDimensionsListExpandAndTop.json
 */
async function managementGroupDimensionsListExpandAndTopLegacy() {
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const resArray = new Array();
  for await (const item of client.dimensions.list(
    "providers/Microsoft.Management/managementGroups/MyMgId",
    { expand: "properties/data", top: 5 },
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
