const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Queries component policy states under subscription scope.
 *
 * @summary Queries component policy states under subscription scope.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2022-04-01/examples/ComponentPolicyStates_QuerySubscriptionScope.json
 */
async function queryLatestComponentPolicyStatesAtSubscriptionScope() {
  const subscriptionId =
    process.env["POLICYINSIGHTS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const componentPolicyStatesResource = "latest";
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const result = await client.componentPolicyStates.listQueryResultsForSubscription(
    subscriptionId,
    componentPolicyStatesResource
  );
  console.log(result);
}
