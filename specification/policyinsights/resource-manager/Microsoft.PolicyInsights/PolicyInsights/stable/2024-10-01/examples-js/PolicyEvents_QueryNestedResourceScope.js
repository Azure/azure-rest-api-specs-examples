const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to queries policy events for the resource.
 *
 * @summary queries policy events for the resource.
 * x-ms-original-file: 2024-10-01/PolicyEvents_QueryNestedResourceScope.json
 */
async function queryAtNestedResourceScope() {
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential);
  const resArray = new Array();
  for await (const item of client.policyEvents.listQueryResultsForResource(
    "default",
    "subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.ServiceFabric/clusters/myCluster/applications/myApplication",
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
