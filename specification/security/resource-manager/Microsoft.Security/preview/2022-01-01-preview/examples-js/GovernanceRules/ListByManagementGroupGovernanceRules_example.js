const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a list of all relevant governance rules over a scope
 *
 * @summary Get a list of all relevant governance rules over a scope
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceRules/ListByManagementGroupGovernanceRules_example.json
 */
async function listGovernanceRulesByManagementGroupScope() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const scope = "providers/Microsoft.Management/managementGroups/contoso";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.governanceRules.list(scope)) {
    resArray.push(item);
  }
  console.log(resArray);
}
