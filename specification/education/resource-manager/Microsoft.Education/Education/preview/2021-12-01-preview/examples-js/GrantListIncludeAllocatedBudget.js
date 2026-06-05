const { EducationManagementClient } = require("@azure/arm-education");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get a list of grants that Microsoft has provided.
 *
 * @summary get a list of grants that Microsoft has provided.
 * x-ms-original-file: 2021-12-01-preview/GrantListIncludeAllocatedBudget.json
 */
async function grantListIncludeAllocatedBudget() {
  const credential = new DefaultAzureCredential();
  const client = new EducationManagementClient(credential);
  const resArray = new Array();
  for await (const item of client.grants.listAll({ includeAllocatedBudget: true })) {
    resArray.push(item);
  }

  console.log(resArray);
}
