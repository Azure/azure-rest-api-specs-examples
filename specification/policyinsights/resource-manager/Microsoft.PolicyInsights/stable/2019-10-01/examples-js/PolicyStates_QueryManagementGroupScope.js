const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Queries policy states for the resources under the management group.
 *
 * @summary Queries policy states for the resources under the management group.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_QueryManagementGroupScope.json
 */
async function queryLatestAtManagementGroupScope() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const policyStatesResource = "latest";
  const managementGroupName = "myManagementGroup";
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.policyStates.listQueryResultsForManagementGroup(
    policyStatesResource,
    managementGroupName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

queryLatestAtManagementGroupScope().catch(console.error);
