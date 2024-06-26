const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get security sub-assessments on all your scanned resources inside a subscription scope
 *
 * @summary Get security sub-assessments on all your scanned resources inside a subscription scope
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/SubAssessments/ListSubscriptionSubAssessments_example.json
 */
async function listSecuritySubAssessments() {
  const scope = "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  const resArray = new Array();
  for await (let item of client.subAssessments.listAll(scope)) {
    resArray.push(item);
  }
  console.log(resArray);
}
