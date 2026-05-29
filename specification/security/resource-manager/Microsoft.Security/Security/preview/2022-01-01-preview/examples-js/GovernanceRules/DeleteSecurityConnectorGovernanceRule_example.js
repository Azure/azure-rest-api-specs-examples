const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete a Governance rule over a given scope
 *
 * @summary delete a Governance rule over a given scope
 * x-ms-original-file: 2022-01-01-preview/GovernanceRules/DeleteSecurityConnectorGovernanceRule_example.json
 */
async function deleteAGovernanceRuleOverSecurityConnectorScope() {
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  await client.governanceRules.delete(
    "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/gcpResourceGroup/providers/Microsoft.Security/securityConnectors/gcpconnector",
    "ad9a8e26-29d9-4829-bb30-e597a58cdbb8",
  );
}
