const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a list of all relevant governanceRules over a subscription level scope
 *
 * @summary Get a list of all relevant governanceRules over a subscription level scope
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceRules/ListBySubscriptionGovernanceRules_example.json
 */
async function listSecurityGovernanceRulesBySubscriptionLevelScope() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.governanceRuleOperations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listSecurityGovernanceRulesBySubscriptionLevelScope().catch(console.error);
