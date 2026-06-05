const { EducationManagementClient } = require("@azure/arm-education");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get a list of labs associated with the provided billing account name and billing profile name.
 *
 * @summary get a list of labs associated with the provided billing account name and billing profile name.
 * x-ms-original-file: 2021-12-01-preview/LabListIncludeBudget.json
 */
async function labListIncludeBudget() {
  const credential = new DefaultAzureCredential();
  const client = new EducationManagementClient(credential);
  const resArray = new Array();
  for await (const item of client.labs.listAll("{billingAccountName}", "{billingProfileName}", {
    includeBudget: true,
  })) {
    resArray.push(item);
  }

  console.log(resArray);
}
