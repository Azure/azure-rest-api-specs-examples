const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a queue authorization rule.
 *
 * @summary Deletes a queue authorization rule.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/Queues/SBQueueAuthorizationRuleDelete.json
 */
async function queueAuthorizationRuleDelete() {
  const subscriptionId =
    process.env["SERVICEBUS_SUBSCRIPTION_ID"] || "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const resourceGroupName = process.env["SERVICEBUS_RESOURCE_GROUP"] || "ArunMonocle";
  const namespaceName = "sdk-namespace-7982";
  const queueName = "sdk-Queues-2317";
  const authorizationRuleName = "sdk-AuthRules-5800";
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const result = await client.queues.deleteAuthorizationRule(
    resourceGroupName,
    namespaceName,
    queueName,
    authorizationRuleName
  );
  console.log(result);
}
