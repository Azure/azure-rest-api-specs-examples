const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a topic authorization rule.
 *
 * @summary Deletes a topic authorization rule.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Topics/SBTopicAuthorizationRuleDelete.json
 */
async function topicAuthorizationRuleDelete() {
  const subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const resourceGroupName = "ArunMonocle";
  const namespaceName = "sdk-Namespace-6261";
  const topicName = "sdk-Topics-1984";
  const authorizationRuleName = "sdk-AuthRules-4310";
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const result = await client.topics.deleteAuthorizationRule(
    resourceGroupName,
    namespaceName,
    topicName,
    authorizationRuleName
  );
  console.log(result);
}

topicAuthorizationRuleDelete().catch(console.error);
