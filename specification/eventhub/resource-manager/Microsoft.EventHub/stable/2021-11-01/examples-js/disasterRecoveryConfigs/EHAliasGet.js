const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves Alias(Disaster Recovery configuration) for primary or secondary namespace
 *
 * @summary Retrieves Alias(Disaster Recovery configuration) for primary or secondary namespace
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/disasterRecoveryConfigs/EHAliasGet.json
 */
async function ehAliasGet() {
  const subscriptionId = "exampleSubscriptionId";
  const resourceGroupName = "exampleResourceGroup";
  const namespaceName = "sdk-Namespace-8859";
  const alias = "sdk-DisasterRecovery-3814";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.disasterRecoveryConfigs.get(resourceGroupName, namespaceName, alias);
  console.log(result);
}

ehAliasGet().catch(console.error);
