const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets lists of resources that supports Privatelinks.
 *
 * @summary Gets lists of resources that supports Privatelinks.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/NameSpaces/PrivateLinkResourcesGet.json
 */
async function nameSpacePrivateLinkResourcesGet() {
  const subscriptionId = "subID";
  const resourceGroupName = "ArunMonocle";
  const namespaceName = "sdk-Namespace-2924";
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const result = await client.privateLinkResources.get(resourceGroupName, namespaceName);
  console.log(result);
}

nameSpacePrivateLinkResourcesGet().catch(console.error);
