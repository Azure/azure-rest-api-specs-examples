const { AzureStackManagementClient } = require("@azure/arm-azurestack");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete the requested Linked Subscription resource.
 *
 * @summary Delete the requested Linked Subscription resource.
 * x-ms-original-file: specification/azurestack/resource-manager/Microsoft.AzureStack/preview/2020-06-01-preview/examples/LinkedSubscription/Delete.json
 */
async function deleteTheRequestedLinkedSubscription() {
  const subscriptionId = "dd8597b4-8739-4467-8b10-f8679f62bfbf";
  const resourceGroup = "azurestack";
  const linkedSubscriptionName = "testlinkedsubscription";
  const credential = new DefaultAzureCredential();
  const client = new AzureStackManagementClient(credential, subscriptionId);
  const result = await client.linkedSubscriptions.delete(resourceGroup, linkedSubscriptionName);
  console.log(result);
}

deleteTheRequestedLinkedSubscription().catch(console.error);
