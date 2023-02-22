const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the primary and secondary connection strings for the namespace.
 *
 * @summary Gets the primary and secondary connection strings for the namespace.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/disasterRecoveryConfigs/SBAliasAuthorizationRuleListKey.json
 */
async function disasterRecoveryConfigsAuthorizationRuleListKey() {
  const subscriptionId = process.env["SERVICEBUS_SUBSCRIPTION_ID"] || "exampleSubscriptionId";
  const resourceGroupName = process.env["SERVICEBUS_RESOURCE_GROUP"] || "exampleResourceGroup";
  const namespaceName = "sdk-Namespace-2702";
  const alias = "sdk-DisasterRecovery-4047";
  const authorizationRuleName = "sdk-Authrules-1746";
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const result = await client.disasterRecoveryConfigs.listKeys(
    resourceGroupName,
    namespaceName,
    alias,
    authorizationRuleName
  );
  console.log(result);
}
