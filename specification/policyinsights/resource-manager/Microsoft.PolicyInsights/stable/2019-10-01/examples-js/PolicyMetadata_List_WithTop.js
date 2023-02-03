const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a list of the policy metadata resources.
 *
 * @summary Get a list of the policy metadata resources.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyMetadata_List_WithTop.json
 */
async function getCollectionOfPolicyMetadataResourcesUsingTopQueryParameter() {
  const subscriptionId =
    process.env["POLICYINSIGHTS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const top = 1;
  const options = { queryOptions: { top: top } };
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.policyMetadataOperations.list(options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
