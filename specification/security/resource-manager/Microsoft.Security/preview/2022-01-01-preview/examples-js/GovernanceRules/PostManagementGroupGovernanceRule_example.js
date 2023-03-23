const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Execute a governance rule
 *
 * @summary Execute a governance rule
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceRules/PostManagementGroupGovernanceRule_example.json
 */
async function executeGovernanceRuleOverManagementGroupScope() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const scope = "providers/Microsoft.Management/managementGroups/contoso";
  const ruleId = "ad9a8e26-29d9-4829-bb30-e597a58cdbb8";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.governanceRules.beginExecuteAndWait(scope, ruleId);
  console.log(result);
}
