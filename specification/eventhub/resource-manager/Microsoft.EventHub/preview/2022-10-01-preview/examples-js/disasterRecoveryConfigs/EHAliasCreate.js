const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a new Alias(Disaster Recovery configuration)
 *
 * @summary Creates or updates a new Alias(Disaster Recovery configuration)
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-10-01-preview/examples/disasterRecoveryConfigs/EHAliasCreate.json
 */
async function ehAliasCreate() {
  const subscriptionId = process.env["EVENTHUB_SUBSCRIPTION_ID"] || "exampleSubscriptionId";
  const resourceGroupName = process.env["EVENTHUB_RESOURCE_GROUP"] || "exampleResourceGroup";
  const namespaceName = "sdk-Namespace-8859";
  const alias = "sdk-DisasterRecovery-3814";
  const parameters = {
    partnerNamespace: "sdk-Namespace-37",
  };
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.disasterRecoveryConfigs.createOrUpdate(
    resourceGroupName,
    namespaceName,
    alias,
    parameters
  );
  console.log(result);
}
