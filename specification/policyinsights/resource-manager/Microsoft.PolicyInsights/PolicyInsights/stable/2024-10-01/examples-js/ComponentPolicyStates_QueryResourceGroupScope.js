const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to queries component policy states under resource group scope.
 *
 * @summary queries component policy states under resource group scope.
 * x-ms-original-file: 2024-10-01/ComponentPolicyStates_QueryResourceGroupScope.json
 */
async function queryLatestComponentPolicyStatesAtResourceGroupScope() {
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential);
  const result = await client.componentPolicyStates.listQueryResultsForResourceGroup(
    "fffedd8f-ffff-fffd-fffd-fffed2f84852",
    "myResourceGroup",
    "latest",
  );
  console.log(result);
}
