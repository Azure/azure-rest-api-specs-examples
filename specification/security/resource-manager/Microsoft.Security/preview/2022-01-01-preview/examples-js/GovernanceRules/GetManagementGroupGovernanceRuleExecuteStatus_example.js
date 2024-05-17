const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get governance rules long run operation result for the requested scope by ruleId and operationId
 *
 * @summary Get governance rules long run operation result for the requested scope by ruleId and operationId
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2022-01-01-preview/examples/GovernanceRules/GetManagementGroupGovernanceRuleExecuteStatus_example.json
 */
async function getGovernanceRulesLongRunOperationResultOverManagementGroup() {
  const scope = "providers/Microsoft.Management/managementGroups/contoso";
  const ruleId = "ad9a8e26-29d9-4829-bb30-e597a58cdbb8";
  const operationId = "58b33f4f-c8c7-4b01-99cc-d437db4d40dd";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  const result = await client.governanceRules.operationResults(scope, ruleId, operationId);
  console.log(result);
}
