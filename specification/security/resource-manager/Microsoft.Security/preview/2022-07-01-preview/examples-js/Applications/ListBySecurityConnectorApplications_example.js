const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a list of all relevant applications over a security connector level scope
 *
 * @summary Get a list of all relevant applications over a security connector level scope
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2022-07-01-preview/examples/Applications/ListBySecurityConnectorApplications_example.json
 */
async function listSecurityApplicationsBySecurityConnectorLevelScope() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const resourceGroupName = process.env["SECURITY_RESOURCE_GROUP"] || "gcpResourceGroup";
  const securityConnectorName = "gcpconnector";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.securityConnectorApplications.list(
    resourceGroupName,
    securityConnectorName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
