const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an AuthorizationRule for a Namespace.
 *
 * @summary Creates or updates an AuthorizationRule for a Namespace.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/NameSpaces/EHNameSpaceAuthorizationRuleCreate.json
 */
async function nameSpaceAuthorizationRuleCreate() {
  const subscriptionId =
    process.env["EVENTHUB_SUBSCRIPTION_ID"] || "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const resourceGroupName = process.env["EVENTHUB_RESOURCE_GROUP"] || "ArunMonocle";
  const namespaceName = "sdk-Namespace-2702";
  const authorizationRuleName = "sdk-Authrules-1746";
  const parameters = { rights: ["Listen", "Send"] };
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.namespaces.createOrUpdateAuthorizationRule(
    resourceGroupName,
    namespaceName,
    authorizationRuleName,
    parameters,
  );
  console.log(result);
}
