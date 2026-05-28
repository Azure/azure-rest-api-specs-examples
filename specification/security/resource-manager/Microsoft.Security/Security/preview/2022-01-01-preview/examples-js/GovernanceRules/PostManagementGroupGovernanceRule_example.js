const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to execute a governance rule
 *
 * @summary execute a governance rule
 * x-ms-original-file: 2022-01-01-preview/GovernanceRules/PostManagementGroupGovernanceRule_example.json
 */
async function executeGovernanceRuleOverManagementGroupScope() {
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  await client.governanceRules.execute(
    "providers/Microsoft.Management/managementGroups/contoso",
    "ad9a8e26-29d9-4829-bb30-e597a58cdbb8",
  );
}
