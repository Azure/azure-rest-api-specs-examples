const { AzureStackManagementClient } = require("@azure/arm-azurestack");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the properties of a Linked Subscription resource.
 *
 * @summary Returns the properties of a Linked Subscription resource.
 * x-ms-original-file: specification/azurestack/resource-manager/Microsoft.AzureStack/preview/2020-06-01-preview/examples/LinkedSubscription/Get.json
 */
async function returnsThePropertiesOfALinkedSubscriptionResource() {
  const subscriptionId = "dd8597b4-8739-4467-8b10-f8679f62bfbf";
  const resourceGroup = "azurestack";
  const linkedSubscriptionName = "testLinkedSubscription";
  const credential = new DefaultAzureCredential();
  const client = new AzureStackManagementClient(credential, subscriptionId);
  const result = await client.linkedSubscriptions.get(resourceGroup, linkedSubscriptionName);
  console.log(result);
}

returnsThePropertiesOfALinkedSubscriptionResource().catch(console.error);
