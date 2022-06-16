const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a topic in the specified namespace.
 *
 * @summary Creates a topic in the specified namespace.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Topics/SBTopicCreate.json
 */
async function topicCreate() {
  const subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const resourceGroupName = "ArunMonocle";
  const namespaceName = "sdk-Namespace-1617";
  const topicName = "sdk-Topics-5488";
  const parameters = { enableExpress: true };
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const result = await client.topics.createOrUpdate(
    resourceGroupName,
    namespaceName,
    topicName,
    parameters
  );
  console.log(result);
}

topicCreate().catch(console.error);
