const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Queries policy events for the resource group level policy assignment.
 *
 * @summary Queries policy events for the resource group level policy assignment.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyEvents_QueryResourceGroupLevelPolicyAssignmentScopeNextLink.json
 */
async function queryAtResourceGroupLevelPolicyAssignmentScopeWithNextLink() {
  const subscriptionId =
    process.env["POLICYINSIGHTS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const policyEventsResource = "default";
  const resourceGroupName = process.env["POLICYINSIGHTS_RESOURCE_GROUP"] || "myResourceGroup";
  const policyAssignmentName = "myPolicyAssignment";
  const skipToken = "WpmWfBSvPhkAK6QD";
  const options = {
    queryOptions: { skipToken: skipToken },
  };
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.policyEvents.listQueryResultsForResourceGroupLevelPolicyAssignment(
    policyEventsResource,
    subscriptionId,
    resourceGroupName,
    policyAssignmentName,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
