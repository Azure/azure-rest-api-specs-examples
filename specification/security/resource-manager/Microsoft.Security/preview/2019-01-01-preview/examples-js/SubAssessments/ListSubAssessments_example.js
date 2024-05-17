const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get security sub-assessments on all your scanned resources inside a scope
 *
 * @summary Get security sub-assessments on all your scanned resources inside a scope
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/SubAssessments/ListSubAssessments_example.json
 */
async function listSecuritySubAssessments() {
  const scope = "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const assessmentName = "82e20e14-edc5-4373-bfc4-f13121257c37";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  const resArray = new Array();
  for await (let item of client.subAssessments.list(scope, assessmentName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
