const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This operation disables the Disaster Recovery and stops replicating changes from primary to secondary namespaces
 *
 * @summary This operation disables the Disaster Recovery and stops replicating changes from primary to secondary namespaces
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-10-01-preview/examples/disasterRecoveryConfigs/EHAliasBreakPairing.json
 */
async function ehAliasBreakPairing() {
  const subscriptionId = process.env["EVENTHUB_SUBSCRIPTION_ID"] || "exampleSubscriptionId";
  const resourceGroupName = process.env["EVENTHUB_RESOURCE_GROUP"] || "exampleResourceGroup";
  const namespaceName = "sdk-Namespace-8859";
  const alias = "sdk-DisasterRecovery-3814";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.disasterRecoveryConfigs.breakPairing(
    resourceGroupName,
    namespaceName,
    alias
  );
  console.log(result);
}
