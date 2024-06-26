const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Queries component policy states for the resource.
 *
 * @summary Queries component policy states for the resource.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2022-04-01/examples/ComponentPolicyStates_QueryResourceScopeExpandPolicyEvaluationDetails.json
 */
async function queryLatestComponentPolicyStatesAtResourceScopeAndExpandPolicyEvaluationDetails() {
  const subscriptionId =
    process.env["POLICYINSIGHTS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceId =
    "subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.ContainerService/managedClusters/myCluster";
  const componentPolicyStatesResource = "latest";
  const filter =
    "componentType eq 'pod' AND componentId eq 'default/test-pod' AND componentName eq 'test-pod'";
  const expand = "PolicyEvaluationDetails";
  const options = {
    filter,
    expand,
  };
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const result = await client.componentPolicyStates.listQueryResultsForResource(
    resourceId,
    componentPolicyStatesResource,
    options
  );
  console.log(result);
}
