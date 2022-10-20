const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Queries policy states for the subscription level policy definition.
 *
 * @summary Queries policy states for the subscription level policy definition.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_QuerySubscriptionLevelPolicyDefinitionScopeNextLink.json
 */
async function queryLatestAtSubscriptionLevelPolicyDefinitionScopeWithNextLink() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const policyStatesResource = "latest";
  const policyDefinitionName = "24813039-7534-408a-9842-eb99f45721b1";
  const skipToken = "WpmWfBSvPhkAK6QD";
  const options = {
    queryOptions: { skipToken: skipToken },
  };
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.policyStates.listQueryResultsForPolicyDefinition(
    policyStatesResource,
    subscriptionId,
    policyDefinitionName,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

queryLatestAtSubscriptionLevelPolicyDefinitionScopeWithNextLink().catch(console.error);
