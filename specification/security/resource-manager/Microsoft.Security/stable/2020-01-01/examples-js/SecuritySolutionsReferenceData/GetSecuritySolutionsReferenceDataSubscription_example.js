const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of all supported Security Solutions for the subscription.
 *
 * @summary Gets a list of all supported Security Solutions for the subscription.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/SecuritySolutionsReferenceData/GetSecuritySolutionsReferenceDataSubscription_example.json
 */
async function getSecuritySolutions() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.securitySolutionsReferenceDataOperations.list();
  console.log(result);
}

getSecuritySolutions().catch(console.error);
