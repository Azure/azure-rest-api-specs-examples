const { EducationManagementClient } = require("@azure/arm-education");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a list of grants that Microsoft has provided.
 *
 * @summary Get a list of grants that Microsoft has provided.
 * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/GrantListIncludeAllocatedBudget.json
 */
async function grantListIncludeAllocatedBudget() {
  const includeAllocatedBudget = true;
  const options = { includeAllocatedBudget };
  const credential = new DefaultAzureCredential();
  const client = new EducationManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.grants.listAll(options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
