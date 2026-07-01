const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to queries component policy states for the resource.
 *
 * @summary queries component policy states for the resource.
 * x-ms-original-file: 2024-10-01/ComponentPolicyStates_QueryResourceScopeExpandPolicyEvaluationDetails.json
 */
async function queryLatestComponentPolicyStatesAtResourceScopeAndExpandPolicyEvaluationDetails() {
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential);
  const result = await client.componentPolicyStates.listQueryResultsForResource(
    "subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.ContainerService/managedClusters/myCluster",
    "latest",
    {
      filter:
        "componentType eq 'pod' AND componentId eq 'default/test-pod' AND componentName eq 'test-pod'",
      expand: "PolicyEvaluationDetails",
    },
  );
  console.log(result);
}
