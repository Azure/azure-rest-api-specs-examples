const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Regenerates primary or secondary connection strings for the topic.
 *
 * @summary Regenerates primary or secondary connection strings for the topic.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/Topics/SBTopicAuthorizationRuleRegenerateKey.json
 */
async function topicAuthorizationRuleRegenerateKey() {
  const subscriptionId =
    process.env["SERVICEBUS_SUBSCRIPTION_ID"] || "e2f361f0-3b27-4503-a9cc-21cfba380093";
  const resourceGroupName = process.env["SERVICEBUS_RESOURCE_GROUP"] || "Default-ServiceBus-WestUS";
  const namespaceName = "sdk-Namespace8408";
  const topicName = "sdk-Topics2075";
  const authorizationRuleName = "sdk-Authrules5067";
  const parameters = { keyType: "PrimaryKey" };
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const result = await client.topics.regenerateKeys(
    resourceGroupName,
    namespaceName,
    topicName,
    authorizationRuleName,
    parameters
  );
  console.log(result);
}
