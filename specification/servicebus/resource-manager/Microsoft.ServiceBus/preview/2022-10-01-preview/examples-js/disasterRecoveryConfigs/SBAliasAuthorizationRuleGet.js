const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an authorization rule for a namespace by rule name.
 *
 * @summary Gets an authorization rule for a namespace by rule name.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/disasterRecoveryConfigs/SBAliasAuthorizationRuleGet.json
 */
async function disasterRecoveryConfigsAuthorizationRuleGet() {
  const subscriptionId = process.env["SERVICEBUS_SUBSCRIPTION_ID"] || "exampleSubscriptionId";
  const resourceGroupName = process.env["SERVICEBUS_RESOURCE_GROUP"] || "exampleResourceGroup";
  const namespaceName = "sdk-Namespace-9080";
  const alias = "sdk-DisasterRecovery-4879";
  const authorizationRuleName = "sdk-Authrules-4879";
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const result = await client.disasterRecoveryConfigs.getAuthorizationRule(
    resourceGroupName,
    namespaceName,
    alias,
    authorizationRuleName
  );
  console.log(result);
}
