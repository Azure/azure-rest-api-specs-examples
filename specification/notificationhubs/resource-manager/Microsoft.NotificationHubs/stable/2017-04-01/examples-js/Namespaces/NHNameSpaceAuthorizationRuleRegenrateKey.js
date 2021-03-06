const { NotificationHubsManagementClient } = require("@azure/arm-notificationhubs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Regenerates the Primary/Secondary Keys to the Namespace Authorization Rule
 *
 * @summary Regenerates the Primary/Secondary Keys to the Namespace Authorization Rule
 * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/Namespaces/NHNameSpaceAuthorizationRuleRegenrateKey.json
 */
async function nameSpaceAuthorizationRuleRegenerateKey() {
  const subscriptionId = "29cfa613-cbbc-4512-b1d6-1b3a92c7fa40";
  const resourceGroupName = "5ktrial";
  const namespaceName = "nh-sdk-ns";
  const authorizationRuleName = "RootManageSharedAccessKey";
  const parameters = { policyKey: "PrimaryKey" };
  const credential = new DefaultAzureCredential();
  const client = new NotificationHubsManagementClient(credential, subscriptionId);
  const result = await client.namespaces.regenerateKeys(
    resourceGroupName,
    namespaceName,
    authorizationRuleName,
    parameters
  );
  console.log(result);
}

nameSpaceAuthorizationRuleRegenerateKey().catch(console.error);
