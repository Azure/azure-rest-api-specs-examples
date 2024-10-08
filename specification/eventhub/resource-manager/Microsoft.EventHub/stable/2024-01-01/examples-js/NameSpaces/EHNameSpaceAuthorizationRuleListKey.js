const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the primary and secondary connection strings for the Namespace.
 *
 * @summary Gets the primary and secondary connection strings for the Namespace.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/NameSpaces/EHNameSpaceAuthorizationRuleListKey.json
 */
async function nameSpaceAuthorizationRuleListKey() {
  const subscriptionId =
    process.env["EVENTHUB_SUBSCRIPTION_ID"] || "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const resourceGroupName = process.env["EVENTHUB_RESOURCE_GROUP"] || "ArunMonocle";
  const namespaceName = "sdk-Namespace-2702";
  const authorizationRuleName = "sdk-Authrules-1746";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.namespaces.listKeys(
    resourceGroupName,
    namespaceName,
    authorizationRuleName,
  );
  console.log(result);
}
