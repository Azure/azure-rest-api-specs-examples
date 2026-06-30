const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to queries component policy states under subscription scope.
 *
 * @summary queries component policy states under subscription scope.
 * x-ms-original-file: 2024-10-01/ComponentPolicyStates_QuerySubscriptionScopeGroupByComponentTypeWithAggregate.json
 */
async function queryLatestComponentPolicyComplianceStateCountGroupedByComponentTypeAtSubscriptionScopeFilteredByGivenAssignment() {
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential);
  const result = await client.componentPolicyStates.listQueryResultsForSubscription(
    "e78961ba-36fe-4739-9212-e3031b4c8db7",
    "latest",
    {
      filter:
        "policyAssignmentId eq '/subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/providers/microsoft.authorization/policyassignments/560050f83dbb4a24974323f8'",
      apply: "groupby((componentType,complianceState),aggregate($count as count))",
    },
  );
  console.log(result);
}
