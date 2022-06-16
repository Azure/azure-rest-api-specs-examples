const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates an authorization rule for a queue.
 *
 * @summary Creates an authorization rule for a queue.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Queues/SBQueueAuthorizationRuleCreate.json
 */
async function queueAuthorizationRuleCreate() {
  const subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const resourceGroupName = "ArunMonocle";
  const namespaceName = "sdk-Namespace-7982";
  const queueName = "sdk-Queues-2317";
  const authorizationRuleName = "sdk-AuthRules-5800";
  const parameters = { rights: ["Listen", "Send"] };
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const result = await client.queues.createOrUpdateAuthorizationRule(
    resourceGroupName,
    namespaceName,
    queueName,
    authorizationRuleName,
    parameters
  );
  console.log(result);
}

queueAuthorizationRuleCreate().catch(console.error);
