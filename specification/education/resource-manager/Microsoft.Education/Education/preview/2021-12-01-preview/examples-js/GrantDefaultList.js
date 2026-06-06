const { EducationManagementClient } = require("@azure/arm-education");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get details for a specific grant linked to the provided billing account and billing profile.
 *
 * @summary get details for a specific grant linked to the provided billing account and billing profile.
 * x-ms-original-file: 2021-12-01-preview/GrantDefaultList.json
 */
async function grantList() {
  const credential = new DefaultAzureCredential();
  const client = new EducationManagementClient(credential);
  const resArray = new Array();
  for await (const item of client.grants.list("{billingAccountName}", "{billingProfileName}", {
    includeAllocatedBudget: false,
  })) {
    resArray.push(item);
  }

  console.log(resArray);
}
