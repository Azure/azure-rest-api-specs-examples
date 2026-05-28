const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get a list of all relevant governance rules over a scope
 *
 * @summary get a list of all relevant governance rules over a scope
 * x-ms-original-file: 2022-01-01-preview/GovernanceRules/ListByManagementGroupGovernanceRules_example.json
 */
async function listGovernanceRulesByManagementGroupScope() {
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  const resArray = new Array();
  for await (const item of client.governanceRules.list(
    "providers/Microsoft.Management/managementGroups/contoso",
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
