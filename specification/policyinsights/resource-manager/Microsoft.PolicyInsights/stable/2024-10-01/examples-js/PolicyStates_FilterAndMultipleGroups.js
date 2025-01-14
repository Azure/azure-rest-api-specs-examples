const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Queries policy states for the resources under the subscription.
 *
 * @summary Queries policy states for the resources under the subscription.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/PolicyStates_FilterAndMultipleGroups.json
 */
async function filterAndMultipleGroups() {
  const policyStatesResource = "latest";
  const subscriptionId = "fffedd8f-ffff-fffd-fffd-fffed2f84852";
  const top = 10;
  const orderBy = "NumNonCompliantResources desc";
  const filter = "IsCompliant eq false";
  const apply =
    "groupby((PolicyAssignmentId, PolicySetDefinitionId, PolicyDefinitionId, PolicyDefinitionReferenceId, ResourceId))/groupby((PolicyAssignmentId, PolicySetDefinitionId, PolicyDefinitionId, PolicyDefinitionReferenceId), aggregate($count as NumNonCompliantResources))";
  const options = {
    top,
    orderBy,
    filter,
    apply,
  };
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential);
  const resArray = new Array();
  for await (let item of client.policyStates.listQueryResultsForSubscription(
    policyStatesResource,
    subscriptionId,
    options,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
