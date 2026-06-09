const { ResourceManagementClient } = require("@azure/arm-resources");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets the entire set of tags on a resource or subscription.
 *
 * @summary gets the entire set of tags on a resource or subscription.
 * x-ms-original-file: 2025-04-01/GetTagsSubscription.json
 */
async function getTagsOnASubscription() {
  const credential = new DefaultAzureCredential();
  const client = new ResourceManagementClient(credential);
  const result = await client.tagsOperations.getAtScope(
    "subscriptions/00000000-0000-0000-0000-000000000000",
  );
  console.log(result);
}
