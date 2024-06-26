const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a Service Bus queue. This operation is idempotent.
 *
 * @summary Creates or updates a Service Bus queue. This operation is idempotent.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/Queues/SBQueueCreate.json
 */
async function queueCreate() {
  const subscriptionId =
    process.env["SERVICEBUS_SUBSCRIPTION_ID"] || "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const resourceGroupName = process.env["SERVICEBUS_RESOURCE_GROUP"] || "ArunMonocle";
  const namespaceName = "sdk-Namespace-3174";
  const queueName = "sdk-Queues-5647";
  const parameters = { enablePartitioning: true };
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const result = await client.queues.createOrUpdate(
    resourceGroupName,
    namespaceName,
    queueName,
    parameters
  );
  console.log(result);
}
