const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a topic subscription.
 *
 * @summary Creates a topic subscription.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/Subscriptions/SBSubscriptionCreate.json
 */
async function subscriptionCreate() {
  const subscriptionId = process.env["SERVICEBUS_SUBSCRIPTION_ID"] || "Subscriptionid";
  const resourceGroupName = process.env["SERVICEBUS_RESOURCE_GROUP"] || "ResourceGroup";
  const namespaceName = "sdk-Namespace-1349";
  const topicName = "sdk-Topics-8740";
  const subscriptionName = "sdk-Subscriptions-2178";
  const parameters = { enableBatchedOperations: true };
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const result = await client.subscriptions.createOrUpdate(
    resourceGroupName,
    namespaceName,
    topicName,
    subscriptionName,
    parameters
  );
  console.log(result);
}
