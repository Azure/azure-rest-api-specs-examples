const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get security assessments on all your scanned resources inside a scope
 *
 * @summary Get security assessments on all your scanned resources inside a scope
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/Assessments/ListAssessments_example.json
 */
async function listSecurityAssessments() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const scope = "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.assessments.list(scope)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listSecurityAssessments().catch(console.error);
