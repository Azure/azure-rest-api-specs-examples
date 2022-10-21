const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Queries policy events for the resource.
 *
 * @summary Queries policy events for the resource.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyEvents_QueryResourceScopeExpandComponents.json
 */
async function queryComponentsPolicyEventsForResourceScopeFilteredByGivenAssignment() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const policyEventsResource = "default";
  const resourceId =
    "subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/Vaults/myKVName";
  const filter =
    "policyAssignmentId eq '/subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/providers/microsoft.authorization/policyassignments/560050f83dbb4a24974323f8'";
  const expand = "components";
  const options = {
    queryOptions: { filter: filter, expand: expand },
  };
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.policyEvents.listQueryResultsForResource(
    policyEventsResource,
    resourceId,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

queryComponentsPolicyEventsForResourceScopeFilteredByGivenAssignment().catch(console.error);
