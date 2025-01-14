const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Queries policy states for the resource.
 *
 * @summary Queries policy states for the resource.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/PolicyStates_QueryNestedResourceScope.json
 */
async function queryAllPolicyStatesAtNestedResourceScope() {
  const policyStatesResource = "default";
  const resourceId =
    "subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.ServiceFabric/clusters/myCluster/applications/myApplication";
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential);
  const resArray = new Array();
  for await (let item of client.policyStates.listQueryResultsForResource(
    policyStatesResource,
    resourceId,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
