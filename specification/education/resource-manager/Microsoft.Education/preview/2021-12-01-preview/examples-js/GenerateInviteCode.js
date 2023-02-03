const { EducationManagementClient } = require("@azure/arm-education");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Generate invite code for a lab
 *
 * @summary Generate invite code for a lab
 * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/GenerateInviteCode.json
 */
async function createLab() {
  const billingAccountName = "{billingAccountName}";
  const billingProfileName = "{billingProfileName}";
  const invoiceSectionName = "{invoiceSectionName}";
  const parameters = { maxStudentCount: 10 };
  const credential = new DefaultAzureCredential();
  const client = new EducationManagementClient(credential);
  const result = await client.labs.generateInviteCode(
    billingAccountName,
    billingProfileName,
    invoiceSectionName,
    parameters
  );
  console.log(result);
}
