const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of external security solutions for the subscription.
 *
 * @summary Gets a list of external security solutions for the subscription.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/ExternalSecuritySolutions/GetExternalSecuritySolutionsSubscription_example.json
 */
async function getExternalSecuritySolutionsOnASubscription() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.externalSecuritySolutions.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
