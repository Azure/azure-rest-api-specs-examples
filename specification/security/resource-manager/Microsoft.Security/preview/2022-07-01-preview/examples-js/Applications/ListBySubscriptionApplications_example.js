const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a list of all relevant applications over a subscription level scope
 *
 * @summary Get a list of all relevant applications over a subscription level scope
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2022-07-01-preview/examples/Applications/ListBySubscriptionApplications_example.json
 */
async function listApplicationsSecurityBySubscriptionLevelScope() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.applications.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
