const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Regenerates the primary or secondary connection strings for the specified Namespace.
 *
 * @summary Regenerates the primary or secondary connection strings for the specified Namespace.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/NameSpaces/EHNameSpaceAuthorizationRuleRegenerateKey.json
 */
async function nameSpaceAuthorizationRuleRegenerateKey() {
  const subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const resourceGroupName = "ArunMonocle";
  const namespaceName = "sdk-Namespace-8980";
  const authorizationRuleName = "sdk-Authrules-8929";
  const parameters = { keyType: "PrimaryKey" };
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.namespaces.regenerateKeys(
    resourceGroupName,
    namespaceName,
    authorizationRuleName,
    parameters
  );
  console.log(result);
}

nameSpaceAuthorizationRuleRegenerateKey().catch(console.error);
