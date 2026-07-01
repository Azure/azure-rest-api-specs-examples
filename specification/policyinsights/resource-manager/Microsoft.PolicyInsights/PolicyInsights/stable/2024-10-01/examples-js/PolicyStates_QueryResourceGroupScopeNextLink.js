const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to queries policy states for the resources under the resource group.
 *
 * @summary queries policy states for the resources under the resource group.
 * x-ms-original-file: 2024-10-01/PolicyStates_QueryResourceGroupScopeNextLink.json
 */
async function queryLatestAtResourceGroupScopeWithNextLink() {
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential);
  const resArray = new Array();
  for await (const item of client.policyStates.listQueryResultsForResourceGroup(
    "latest",
    "fffedd8f-ffff-fffd-fffd-fffed2f84852",
    "myResourceGroup",
    { queryOptions: { skipToken: "WpmWfBSvPhkAK6QD" } },
  )) {
    resArray.push(item);
  }

  console.log(resArray);
}
