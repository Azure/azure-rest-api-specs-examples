const { ResourceManagementClient } = require("@azure/arm-resources");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Deletes the entire set of tags on a resource or subscription.
 *
 * @summary Deletes the entire set of tags on a resource or subscription.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2025-03-01/examples/DeleteTagsSubscription.json
 */
async function updateTagsOnASubscription() {
  const scope = "subscriptions/00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new ResourceManagementClient(credential);
  const result = await client.tagsOperations.beginDeleteAtScopeAndWait(scope);
  console.log(result);
}
