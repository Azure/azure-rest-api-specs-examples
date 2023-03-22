const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of discovered Security Solutions for the subscription and location.
 *
 * @summary Gets a list of discovered Security Solutions for the subscription and location.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/DiscoveredSecuritySolutions/GetDiscoveredSecuritySolutionsSubscriptionLocation_example.json
 */
async function getDiscoveredSecuritySolutionsFromASecurityDataLocation() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const ascLocation = "centralus";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.discoveredSecuritySolutions.listByHomeRegion(ascLocation)) {
    resArray.push(item);
  }
  console.log(resArray);
}
