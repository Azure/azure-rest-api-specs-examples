const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Queries policy events for the resources under the resource group.
 *
 * @summary Queries policy events for the resources under the resource group.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyEvents_QueryResourceGroupScope.json
 */
async function queryAtResourceGroupScope() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const policyEventsResource = "default";
  const resourceGroupName = "myResourceGroup";
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.policyEvents.listQueryResultsForResourceGroup(
    policyEventsResource,
    subscriptionId,
    resourceGroupName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

queryAtResourceGroupScope().catch(console.error);
