const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to For a specified subscription, list the available security controls, their assessments, and the max score
 *
 * @summary For a specified subscription, list the available security controls, their assessments, and the max score
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/secureScoreControlDefinitions/ListSecureScoreControlDefinitions_subscription_example.json
 */
async function listSecurityControlsDefinitionBySubscription() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.secureScoreControlDefinitions.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
