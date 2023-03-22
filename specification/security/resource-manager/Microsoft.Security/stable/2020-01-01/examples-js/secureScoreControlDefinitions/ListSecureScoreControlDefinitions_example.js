const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List the available security controls, their assessments, and the max score
 *
 * @summary List the available security controls, their assessments, and the max score
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/secureScoreControlDefinitions/ListSecureScoreControlDefinitions_example.json
 */
async function listSecurityControlsDefinition() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.secureScoreControlDefinitions.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
