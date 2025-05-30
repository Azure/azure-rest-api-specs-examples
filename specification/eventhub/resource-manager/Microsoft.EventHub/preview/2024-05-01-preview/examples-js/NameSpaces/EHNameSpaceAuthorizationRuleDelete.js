const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Deletes an AuthorizationRule for a Namespace.
 *
 * @summary Deletes an AuthorizationRule for a Namespace.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/preview/2024-05-01-preview/examples/NameSpaces/EHNameSpaceAuthorizationRuleDelete.json
 */
async function nameSpaceAuthorizationRuleDelete() {
  const subscriptionId =
    process.env["EVENTHUB_SUBSCRIPTION_ID"] || "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const resourceGroupName = process.env["EVENTHUB_RESOURCE_GROUP"] || "ArunMonocle";
  const namespaceName = "sdk-Namespace-8980";
  const authorizationRuleName = "sdk-Authrules-8929";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.namespaces.deleteAuthorizationRule(
    resourceGroupName,
    namespaceName,
    authorizationRuleName,
  );
  console.log(result);
}
