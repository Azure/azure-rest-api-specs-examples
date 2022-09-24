const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a specific governanceRule execution status for the requested scope by ruleId and operationId
 *
 * @summary Get a specific governanceRule execution status for the requested scope by ruleId and operationId
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceRules/GetSecurityConnectorGovernanceRuleExecuteStatus_example.json
 */
async function getSecurityGovernanceRulesExecutionStatusBySpecificGovernanceRuleId() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const resourceGroupName = "gcpResourceGroup";
  const securityConnectorName = "gcpconnector";
  const ruleId = "ad9a8e26-29d9-4829-bb30-e597a58cdbb8";
  const operationId = "58b33f4f-c8c7-4b01-99cc-d437db4d40dd";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.securityConnectorGovernanceRulesExecuteStatus.beginGetAndWait(
    resourceGroupName,
    securityConnectorName,
    ruleId,
    operationId
  );
  console.log(result);
}

getSecurityGovernanceRulesExecutionStatusBySpecificGovernanceRuleId().catch(console.error);
