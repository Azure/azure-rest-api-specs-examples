const { EducationManagementClient } = require("@azure/arm-education");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get details for a specific grant linked to the provided billing account and billing profile.
 *
 * @summary Get details for a specific grant linked to the provided billing account and billing profile.
 * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/GrantDefaultListIncludeAllocatedBudget.json
 */
async function grantListIncludeAllocatedBudget() {
  const billingAccountName = "{billingAccountName}";
  const billingProfileName = "{billingProfileName}";
  const includeAllocatedBudget = false;
  const options = { includeAllocatedBudget };
  const credential = new DefaultAzureCredential();
  const client = new EducationManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.grants.list(billingAccountName, billingProfileName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
