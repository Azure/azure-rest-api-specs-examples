const { EducationManagementClient } = require("@azure/arm-education");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a list of labs associated with the provided billing account name and billing profile name.
 *
 * @summary Get a list of labs associated with the provided billing account name and billing profile name.
 * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/LabList.json
 */
async function labList() {
  const billingAccountName = "{billingAccountName}";
  const billingProfileName = "{billingProfileName}";
  const includeBudget = false;
  const options = { includeBudget };
  const credential = new DefaultAzureCredential();
  const client = new EducationManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.labs.listAll(billingAccountName, billingProfileName, options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
