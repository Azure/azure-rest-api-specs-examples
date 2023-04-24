const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets lists of resources that supports Privatelinks.
 *
 * @summary Gets lists of resources that supports Privatelinks.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-10-01-preview/examples/NameSpaces/PrivateLinkResourcesGet.json
 */
async function nameSpacePrivateLinkResourcesGet() {
  const subscriptionId = process.env["EVENTHUB_SUBSCRIPTION_ID"] || "subID";
  const resourceGroupName = process.env["EVENTHUB_RESOURCE_GROUP"] || "ArunMonocle";
  const namespaceName = "sdk-Namespace-2924";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.privateLinkResources.get(resourceGroupName, namespaceName);
  console.log(result);
}
