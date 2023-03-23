const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get metadata information on all assessment types in a specific subscription
 *
 * @summary Get metadata information on all assessment types in a specific subscription
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2021-06-01/examples/AssessmentsMetadata/ListAssessmentsMetadata_subscription_example.json
 */
async function listSecurityAssessmentMetadataForSubscription() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "0980887d-03d6-408c-9566-532f3456804e";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.assessmentsMetadata.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
