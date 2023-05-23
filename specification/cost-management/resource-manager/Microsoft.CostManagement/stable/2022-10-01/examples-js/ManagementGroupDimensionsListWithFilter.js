const { CostManagementClient } = require("@azure/arm-costmanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the dimensions by the defined scope.
 *
 * @summary Lists the dimensions by the defined scope.
 * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ManagementGroupDimensionsListWithFilter.json
 */
async function managementGroupDimensionsListWithFilterLegacy() {
  const scope = "providers/Microsoft.Management/managementGroups/MyMgId";
  const filter = "properties/category eq 'resourceId'";
  const expand = "properties/data";
  const top = 5;
  const options = { filter, expand, top };
  const credential = new DefaultAzureCredential();
  const client = new CostManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.dimensions.list(scope, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
