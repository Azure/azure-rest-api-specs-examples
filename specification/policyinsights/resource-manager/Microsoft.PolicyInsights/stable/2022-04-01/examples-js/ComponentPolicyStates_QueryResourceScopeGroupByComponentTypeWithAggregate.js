const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Queries component policy states for the resource.
 *
 * @summary Queries component policy states for the resource.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2022-04-01/examples/ComponentPolicyStates_QueryResourceScopeGroupByComponentTypeWithAggregate.json
 */
async function queryLatestComponentPolicyComplianceStateCountGroupedByComponentTypeAtResourceScopeFilteredByGivenAssignment() {
  const subscriptionId =
    process.env["POLICYINSIGHTS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceId =
    "subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/Vaults/myKVName";
  const componentPolicyStatesResource = "latest";
  const filter =
    "policyAssignmentId eq '/subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/providers/microsoft.authorization/policyassignments/560050f83dbb4a24974323f8'";
  const apply = "groupby((componentType,complianceState),aggregate($count as count))";
  const options = {
    filter,
    apply,
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
