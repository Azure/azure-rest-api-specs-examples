const { EducationManagementClient } = require("@azure/arm-education");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete the specified student based on the student alias.
 *
 * @summary Delete the specified student based on the student alias.
 * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/DeleteStudent.json
 */
async function deleteLab() {
  const billingAccountName = "{billingAccountName}";
  const billingProfileName = "{billingProfileName}";
  const invoiceSectionName = "{invoiceSectionName}";
  const studentAlias = "{studentAlias}";
  const credential = new DefaultAzureCredential();
  const client = new EducationManagementClient(credential);
  const result = await client.students.delete(
    billingAccountName,
    billingProfileName,
    invoiceSectionName,
    studentAlias
  );
  console.log(result);
}
