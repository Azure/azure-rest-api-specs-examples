const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Security compliance results in the subscription
 *
 * @summary Security compliance results in the subscription
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2017-08-01/examples/ComplianceResults/ListComplianceResults_example.json
 */
async function getComplianceResultsOnSubscription() {
  const scope = "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  const resArray = new Array();
  for await (let item of client.complianceResults.list(scope)) {
    resArray.push(item);
  }
  console.log(resArray);
}
