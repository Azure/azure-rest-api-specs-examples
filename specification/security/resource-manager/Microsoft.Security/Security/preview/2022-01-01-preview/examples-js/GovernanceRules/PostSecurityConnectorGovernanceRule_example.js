const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to execute a governance rule
 *
 * @summary execute a governance rule
 * x-ms-original-file: 2022-01-01-preview/GovernanceRules/PostSecurityConnectorGovernanceRule_example.json
 */
async function executeGovernanceRuleOverSecurityConnectorScope() {
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  await client.governanceRules.execute(
    "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/gcpResourceGroup/providers/Microsoft.Security/securityConnectors/gcpconnector",
    "ad9a8e26-29d9-4829-bb30-e597a58cdbb8",
  );
}
