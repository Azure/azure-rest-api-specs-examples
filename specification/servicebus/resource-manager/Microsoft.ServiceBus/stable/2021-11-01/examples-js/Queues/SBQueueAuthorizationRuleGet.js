const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an authorization rule for a queue by rule name.
 *
 * @summary Gets an authorization rule for a queue by rule name.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Queues/SBQueueAuthorizationRuleGet.json
 */
async function queueAuthorizationRuleGet() {
  const subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const resourceGroupName = "ArunMonocle";
  const namespaceName = "sdk-Namespace-7982";
  const queueName = "sdk-Queues-2317";
  const authorizationRuleName = "sdk-AuthRules-5800";
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const result = await client.queues.getAuthorizationRule(
    resourceGroupName,
    namespaceName,
    queueName,
    authorizationRuleName
  );
  console.log(result);
}

queueAuthorizationRuleGet().catch(console.error);
