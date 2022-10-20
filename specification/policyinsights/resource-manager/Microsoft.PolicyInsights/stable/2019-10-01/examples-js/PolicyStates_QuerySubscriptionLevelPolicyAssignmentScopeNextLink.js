const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Queries policy states for the subscription level policy assignment.
 *
 * @summary Queries policy states for the subscription level policy assignment.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_QuerySubscriptionLevelPolicyAssignmentScopeNextLink.json
 */
async function queryLatestAtSubscriptionLevelPolicyAssignmentScopeWithNextLink() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const policyStatesResource = "latest";
  const policyAssignmentName = "ec8f9645-8ecb-4abb-9c0b-5292f19d4003";
  const skipToken = "WpmWfBSvPhkAK6QD";
  const options = {
    queryOptions: { skipToken: skipToken },
  };
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.policyStates.listQueryResultsForSubscriptionLevelPolicyAssignment(
    policyStatesResource,
    subscriptionId,
    policyAssignmentName,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

queryLatestAtSubscriptionLevelPolicyAssignmentScopeWithNextLink().catch(console.error);
