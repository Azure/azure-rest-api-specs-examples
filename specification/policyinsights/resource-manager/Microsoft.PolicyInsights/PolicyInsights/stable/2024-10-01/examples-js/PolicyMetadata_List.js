const { PolicyInsightsClient } = require("@azure/arm-policyinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get a list of the policy metadata resources.
 *
 * @summary get a list of the policy metadata resources.
 * x-ms-original-file: 2024-10-01/PolicyMetadata_List.json
 */
async function getCollectionOfPolicyMetadataResources() {
  const credential = new DefaultAzureCredential();
  const client = new PolicyInsightsClient(credential);
  const resArray = new Array();
  for await (const item of client.policyMetadataOperations.list()) {
    resArray.push(item);
  }

  console.log(resArray);
}
