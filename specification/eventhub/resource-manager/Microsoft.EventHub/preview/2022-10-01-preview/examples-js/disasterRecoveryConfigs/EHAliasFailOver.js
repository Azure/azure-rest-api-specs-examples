const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Invokes GEO DR failover and reconfigure the alias to point to the secondary namespace
 *
 * @summary Invokes GEO DR failover and reconfigure the alias to point to the secondary namespace
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-10-01-preview/examples/disasterRecoveryConfigs/EHAliasFailOver.json
 */
async function ehAliasFailOver() {
  const subscriptionId = process.env["EVENTHUB_SUBSCRIPTION_ID"] || "exampleSubscriptionId";
  const resourceGroupName = process.env["EVENTHUB_RESOURCE_GROUP"] || "exampleResourceGroup";
  const namespaceName = "sdk-Namespace-8859";
  const alias = "sdk-DisasterRecovery-3814";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.disasterRecoveryConfigs.failOver(
    resourceGroupName,
    namespaceName,
    alias
  );
  console.log(result);
}
