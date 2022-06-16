const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the topics in a namespace.
 *
 * @summary Gets all the topics in a namespace.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Topics/SBTopicListByNameSpace.json
 */
async function topicGet() {
  const subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const resourceGroupName = "Default-ServiceBus-WestUS";
  const namespaceName = "sdk-Namespace-1617";
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.topics.listByNamespace(resourceGroupName, namespaceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

topicGet().catch(console.error);
