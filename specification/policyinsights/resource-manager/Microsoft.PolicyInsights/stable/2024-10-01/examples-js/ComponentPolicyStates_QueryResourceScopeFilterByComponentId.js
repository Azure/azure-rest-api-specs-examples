const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Queries component policy states for the resource.
 *
 * @summary Queries component policy states for the resource.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/ComponentPolicyStates_QueryResourceScopeFilterByComponentId.json
 */
async function queryLatestComponentPolicyComplianceStateAtResourceScopeFilteredByGivenComponentId() {
  const resourceId =
    "subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/Vaults/myKVName";
  const componentPolicyStatesResource = "latest";
  const filter = "componentId eq cert-RSA-cert-3";
  const options = { filter };
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential);
  const result = await client.componentPolicyStates.listQueryResultsForResource(
    resourceId,
    componentPolicyStatesResource,
    options,
  );
  console.log(result);
}
