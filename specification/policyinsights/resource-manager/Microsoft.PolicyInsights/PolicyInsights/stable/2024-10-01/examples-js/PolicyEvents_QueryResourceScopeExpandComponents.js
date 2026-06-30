const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to queries policy events for the resource.
 *
 * @summary queries policy events for the resource.
 * x-ms-original-file: 2024-10-01/PolicyEvents_QueryResourceScopeExpandComponents.json
 */
async function queryComponentsPolicyEventsForResourceScopeFilteredByGivenAssignment() {
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential);
  const resArray = new Array();
  for await (const item of client.policyEvents.listQueryResultsForResource(
    "default",
    "subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/Vaults/myKVName",
    {
      queryOptions: {
        filter:
          "policyAssignmentId eq '/subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/providers/microsoft.authorization/policyassignments/560050f83dbb4a24974323f8'",
        expand: "components",
      },
    },
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
