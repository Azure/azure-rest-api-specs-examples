const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Queries policy tracked resources under the resource group.
 *
 * @summary Queries policy tracked resources under the resource group.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/preview/2018-07-01-preview/examples/PolicyTrackedResources_QueryResourceGroupScopeWithFilterAndTop.json
 */
async function queryAtResourceGroupScopeUsingQueryParameters() {
  const subscriptionId =
    process.env["POLICYINSIGHTS_SUBSCRIPTION_ID"] || "fffedd8f-ffff-fffd-fffd-fffed2f84852";
  const resourceGroupName = process.env["POLICYINSIGHTS_RESOURCE_GROUP"] || "myResourceGroup";
  const policyTrackedResourcesResource = "default";
  const top = 1;
  const filter =
    "PolicyAssignmentId eq '/subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/resourceGroups/myResourceGroup/providers/Microsoft.Authorization/policyAssignments/myPolicyAssignment' AND TrackedResourceId eq '/subscriptions/fff8dfdb-fff3-fff0-fff4-fffdcbe6b2ef/resourceGroups/myResourceGroup/providers/Microsoft.Example/exampleResourceType/myResource/nestedResourceType/TrackedResource1'";
  const options = { top, filter };
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.policyTrackedResources.listQueryResultsForResourceGroup(
    resourceGroupName,
    policyTrackedResourcesResource,
    options,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
