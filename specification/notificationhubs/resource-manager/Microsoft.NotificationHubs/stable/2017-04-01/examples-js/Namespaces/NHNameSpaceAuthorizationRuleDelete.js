const { NotificationHubsManagementClient } = require("@azure/arm-notificationhubs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a namespace authorization rule
 *
 * @summary Deletes a namespace authorization rule
 * x-ms-original-file: specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/Namespaces/NHNameSpaceAuthorizationRuleDelete.json
 */
async function nameSpaceAuthorizationRuleDelete() {
  const subscriptionId = "29cfa613-cbbc-4512-b1d6-1b3a92c7fa40";
  const resourceGroupName = "5ktrial";
  const namespaceName = "nh-sdk-ns";
  const authorizationRuleName = "RootManageSharedAccessKey";
  const credential = new DefaultAzureCredential();
  const client = new NotificationHubsManagementClient(credential, subscriptionId);
  const result = await client.namespaces.deleteAuthorizationRule(
    resourceGroupName,
    namespaceName,
    authorizationRuleName
  );
  console.log(result);
}

nameSpaceAuthorizationRuleDelete().catch(console.error);
