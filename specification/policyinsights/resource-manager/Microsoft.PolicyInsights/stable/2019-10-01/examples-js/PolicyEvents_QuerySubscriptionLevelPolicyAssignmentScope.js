const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Queries policy events for the subscription level policy assignment.
 *
 * @summary Queries policy events for the subscription level policy assignment.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyEvents_QuerySubscriptionLevelPolicyAssignmentScope.json
 */
async function queryAtSubscriptionLevelPolicyAssignmentScope() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const policyEventsResource = "default";
  const policyAssignmentName = "ec8f9645-8ecb-4abb-9c0b-5292f19d4003";
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.policyEvents.listQueryResultsForSubscriptionLevelPolicyAssignment(
    policyEventsResource,
    subscriptionId,
    policyAssignmentName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

queryAtSubscriptionLevelPolicyAssignmentScope().catch(console.error);
