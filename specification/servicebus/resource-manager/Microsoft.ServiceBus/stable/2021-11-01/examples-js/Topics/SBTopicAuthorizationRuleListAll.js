const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets authorization rules for a topic.
 *
 * @summary Gets authorization rules for a topic.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Topics/SBTopicAuthorizationRuleListAll.json
 */
async function topicAuthorizationRuleListAll() {
  const subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const resourceGroupName = "ArunMonocle";
  const namespaceName = "sdk-Namespace-6261";
  const topicName = "sdk-Topics-1984";
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.topics.listAuthorizationRules(
    resourceGroupName,
    namespaceName,
    topicName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

topicAuthorizationRuleListAll().catch(console.error);
