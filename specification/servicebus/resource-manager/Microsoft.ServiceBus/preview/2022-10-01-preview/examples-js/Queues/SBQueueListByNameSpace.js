const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the queues within a namespace.
 *
 * @summary Gets the queues within a namespace.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/Queues/SBQueueListByNameSpace.json
 */
async function queueListByNameSpace() {
  const subscriptionId =
    process.env["SERVICEBUS_SUBSCRIPTION_ID"] || "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const resourceGroupName = process.env["SERVICEBUS_RESOURCE_GROUP"] || "ArunMonocle";
  const namespaceName = "sdk-Namespace-3174";
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.queues.listByNamespace(resourceGroupName, namespaceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
