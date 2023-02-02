const { EducationManagementClient } = require("@azure/arm-education");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a list of details about students that are associated with the specified lab.
 *
 * @summary Get a list of details about students that are associated with the specified lab.
 * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/StudentList.json
 */
async function studentList() {
  const billingAccountName = "{billingAccountName}";
  const billingProfileName = "{billingProfileName}";
  const invoiceSectionName = "{invoiceSectionName}";
  const credential = new DefaultAzureCredential();
  const client = new EducationManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.students.list(
    billingAccountName,
    billingProfileName,
    invoiceSectionName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
