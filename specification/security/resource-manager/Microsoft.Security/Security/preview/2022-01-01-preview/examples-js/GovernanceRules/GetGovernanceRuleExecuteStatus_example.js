const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get governance rules long run operation result for the requested scope by ruleId and operationId
 *
 * @summary get governance rules long run operation result for the requested scope by ruleId and operationId
 * x-ms-original-file: 2022-01-01-preview/GovernanceRules/GetGovernanceRuleExecuteStatus_example.json
 */
async function getGovernanceRulesLongRunOperationResultOverSubscription() {
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  const result = await client.governanceRules.operationResults(
    "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23",
    "ad9a8e26-29d9-4829-bb30-e597a58cdbb8",
    "58b33f4f-c8c7-4b01-99cc-d437db4d40dd",
  );
  console.log(result);
}
