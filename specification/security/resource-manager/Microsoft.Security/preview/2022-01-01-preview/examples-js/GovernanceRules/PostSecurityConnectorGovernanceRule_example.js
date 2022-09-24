const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Execute a security GovernanceRule on the given security connector.
 *
 * @summary Execute a security GovernanceRule on the given security connector.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceRules/PostSecurityConnectorGovernanceRule_example.json
 */
async function executeGovernanceRule() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const resourceGroupName = "gcpResourceGroup";
  const securityConnectorName = "gcpconnector";
  const ruleId = "ad9a8e26-29d9-4829-bb30-e597a58cdbb8";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.governanceRules.beginRuleIdExecuteSingleSecurityConnectorAndWait(
    resourceGroupName,
    securityConnectorName,
    ruleId
  );
  console.log(result);
}

executeGovernanceRule().catch(console.error);
