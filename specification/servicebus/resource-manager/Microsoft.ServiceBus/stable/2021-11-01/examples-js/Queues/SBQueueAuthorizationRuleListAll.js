const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all authorization rules for a queue.
 *
 * @summary Gets all authorization rules for a queue.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Queues/SBQueueAuthorizationRuleListAll.json
 */
async function queueAuthorizationRuleListAll() {
  const subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const resourceGroupName = "ArunMonocle";
  const namespaceName = "sdk-Namespace-7982";
  const queueName = "sdk-Queues-2317";
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.queues.listAuthorizationRules(
    resourceGroupName,
    namespaceName,
    queueName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

queueAuthorizationRuleListAll().catch(console.error);
