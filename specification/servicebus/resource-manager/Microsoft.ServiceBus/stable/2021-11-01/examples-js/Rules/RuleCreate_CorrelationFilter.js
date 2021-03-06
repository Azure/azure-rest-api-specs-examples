const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new rule and updates an existing rule
 *
 * @summary Creates a new rule and updates an existing rule
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Rules/RuleCreate_CorrelationFilter.json
 */
async function rulesCreateCorrelationFilter() {
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourceGroupName";
  const namespaceName = "sdk-Namespace-1319";
  const topicName = "sdk-Topics-2081";
  const subscriptionName = "sdk-Subscriptions-8691";
  const ruleName = "sdk-Rules-6571";
  const parameters = {
    correlationFilter: { properties: { topicHint: "Crop" } },
    filterType: "CorrelationFilter",
  };
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const result = await client.rules.createOrUpdate(
    resourceGroupName,
    namespaceName,
    topicName,
    subscriptionName,
    ruleName,
    parameters
  );
  console.log(result);
}

rulesCreateCorrelationFilter().catch(console.error);
