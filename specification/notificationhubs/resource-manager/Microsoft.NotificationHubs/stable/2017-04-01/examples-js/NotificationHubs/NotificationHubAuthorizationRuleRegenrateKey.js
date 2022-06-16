const { NotificationHubsManagementClient } = require("@azure/arm-notificationhubs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Regenerates the Primary/Secondary Keys to the NotificationHub Authorization Rule
 *
 * @summary Regenerates the Primary/Secondary Keys to the NotificationHub Authorization Rule
 * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/NotificationHubs/NotificationHubAuthorizationRuleRegenrateKey.json
 */
async function notificationHubAuthorizationRuleRegenrateKey() {
  const subscriptionId = "29cfa613-cbbc-4512-b1d6-1b3a92c7fa40";
  const resourceGroupName = "5ktrial";
  const namespaceName = "nh-sdk-ns";
  const notificationHubName = "nh-sdk-hub";
  const authorizationRuleName = "DefaultListenSharedAccessSignature";
  const parameters = { policyKey: "PrimaryKey" };
  const credential = new DefaultAzureCredential();
  const client = new NotificationHubsManagementClient(credential, subscriptionId);
  const result = await client.notificationHubs.regenerateKeys(
    resourceGroupName,
    namespaceName,
    notificationHubName,
    authorizationRuleName,
    parameters
  );
  console.log(result);
}

notificationHubAuthorizationRuleRegenrateKey().catch(console.error);
