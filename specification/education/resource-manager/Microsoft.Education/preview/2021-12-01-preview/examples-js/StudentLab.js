const { EducationManagementClient } = require("@azure/arm-education");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the details for a specified lab associated with the student lab.
 *
 * @summary Get the details for a specified lab associated with the student lab.
 * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/StudentLab.json
 */
async function studentLab() {
  const studentLabName = "{studentLabName}";
  const credential = new DefaultAzureCredential();
  const client = new EducationManagementClient(credential);
  const result = await client.studentLabs.get(studentLabName);
  console.log(result);
}
