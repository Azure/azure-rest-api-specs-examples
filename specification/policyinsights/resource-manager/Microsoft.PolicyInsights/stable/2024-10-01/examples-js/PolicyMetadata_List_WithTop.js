const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Get a list of the policy metadata resources.
 *
 * @summary Get a list of the policy metadata resources.
 * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/PolicyMetadata_List_WithTop.json
 */
async function getCollectionOfPolicyMetadataResourcesUsingTopQueryParameter() {
  const top = 1;
  const options = { top };
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential);
  const resArray = new Array();
  for await (let item of client.policyMetadataOperations.list(options)) {
    resArray.push(item);
  }
  console.log(resArray);
}
