const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the primary and secondary connection strings for the Namespace.
 *
 * @summary Gets the primary and secondary connection strings for the Namespace.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/disasterRecoveryConfigs/EHAliasAuthorizationRuleListKey.json
 */
async function nameSpaceAuthorizationRuleListKey() {
  const subscriptionId = process.env["EVENTHUB_SUBSCRIPTION_ID"] || "exampleSubscriptionId";
  const resourceGroupName = process.env["EVENTHUB_RESOURCE_GROUP"] || "exampleResourceGroup";
  const namespaceName = "sdk-Namespace-2702";
  const alias = "sdk-DisasterRecovery-4047";
  const authorizationRuleName = "sdk-Authrules-1746";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.disasterRecoveryConfigs.listKeys(
    resourceGroupName,
    namespaceName,
    alias,
    authorizationRuleName,
  );
  console.log(result);
}
