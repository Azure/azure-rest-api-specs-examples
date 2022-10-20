const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Queries policy states for the subscription level policy set definition.
 *
 * @summary Queries policy states for the subscription level policy set definition.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_QuerySubscriptionLevelPolicySetDefinitionScopeNextLink.json
 */
async function queryLatestAtSubscriptionLevelPolicySetDefinitionScopeWithNextLink() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const policyStatesResource = "latest";
  const policySetDefinitionName = "3e3807c1-65c9-49e0-a406-82d8ae3e338c";
  const skipToken = "WpmWfBSvPhkAK6QD";
  const options = {
    queryOptions: { skipToken: skipToken },
  };
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.policyStates.listQueryResultsForPolicySetDefinition(
    policyStatesResource,
    subscriptionId,
    policySetDefinitionName,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

queryLatestAtSubscriptionLevelPolicySetDefinitionScopeWithNextLink().catch(console.error);
