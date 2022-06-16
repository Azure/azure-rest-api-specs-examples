const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves the description for the specified rule.
 *
 * @summary Retrieves the description for the specified rule.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Rules/RuleGet.json
 */
async function rulesGet() {
  const subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const resourceGroupName = "ArunMonocle";
  const namespaceName = "sdk-Namespace-1319";
  const topicName = "sdk-Topics-2081";
  const subscriptionName = "sdk-Subscriptions-8691";
  const ruleName = "sdk-Rules-6571";
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const result = await client.rules.get(
    resourceGroupName,
    namespaceName,
    topicName,
    subscriptionName,
    ruleName
  );
  console.log(result);
}

rulesGet().catch(console.error);
