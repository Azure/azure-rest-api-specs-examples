const { ResourceManagementClient } = require("@azure/arm-resources-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the entire set of tags on a resource or subscription.
 *
 * @summary Gets the entire set of tags on a resource or subscription.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2019-10-01/examples/GetTagsSubscription.json
 */
async function getTagsOnASubscription() {
  const subscriptionId =
    process.env["RESOURCES_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const scope = "subscriptions/eaee6a92-e973-4922-9471-3a0a6abf81cd";
  const credential = new DefaultAzureCredential();
  const client = new ResourceManagementClient(credential, subscriptionId);
  const result = await client.tagsOperations.getAtScope(scope);
  console.log(result);
}
