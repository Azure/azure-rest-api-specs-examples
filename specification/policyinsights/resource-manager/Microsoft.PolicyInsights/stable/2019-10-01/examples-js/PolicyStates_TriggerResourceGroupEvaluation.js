const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Triggers a policy evaluation scan for all the resources under the resource group.
 *
 * @summary Triggers a policy evaluation scan for all the resources under the resource group.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_TriggerResourceGroupEvaluation.json
 */
async function triggerEvaluationsForAllResourcesInAResourceGroup() {
  const subscriptionId =
    process.env["POLICYINSIGHTS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["POLICYINSIGHTS_RESOURCE_GROUP"] || "myResourceGroup";
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const result = await client.policyStates.beginTriggerResourceGroupEvaluationAndWait(
    subscriptionId,
    resourceGroupName
  );
  console.log(result);
}
