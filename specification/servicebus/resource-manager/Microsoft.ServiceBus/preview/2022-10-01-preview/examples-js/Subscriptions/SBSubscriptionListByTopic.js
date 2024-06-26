const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all the subscriptions under a specified topic.
 *
 * @summary List all the subscriptions under a specified topic.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/Subscriptions/SBSubscriptionListByTopic.json
 */
async function subscriptionListByTopic() {
  const subscriptionId = process.env["SERVICEBUS_SUBSCRIPTION_ID"] || "5{Subscriptionid}";
  const resourceGroupName = process.env["SERVICEBUS_RESOURCE_GROUP"] || "ResourceGroup";
  const namespaceName = "sdk-Namespace-1349";
  const topicName = "sdk-Topics-8740";
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.subscriptions.listByTopic(
    resourceGroupName,
    namespaceName,
    topicName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
