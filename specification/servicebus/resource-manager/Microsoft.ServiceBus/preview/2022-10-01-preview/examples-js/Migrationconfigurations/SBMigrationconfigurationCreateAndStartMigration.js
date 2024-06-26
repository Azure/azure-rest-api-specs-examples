const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates Migration configuration and starts migration of entities from Standard to Premium namespace
 *
 * @summary Creates Migration configuration and starts migration of entities from Standard to Premium namespace
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/Migrationconfigurations/SBMigrationconfigurationCreateAndStartMigration.json
 */
async function migrationConfigurationsStartMigration() {
  const subscriptionId = process.env["SERVICEBUS_SUBSCRIPTION_ID"] || "SubscriptionId";
  const resourceGroupName = process.env["SERVICEBUS_RESOURCE_GROUP"] || "ResourceGroup";
  const namespaceName = "sdk-Namespace-41";
  const configName = "$default";
  const parameters = {
    postMigrationName: "sdk-PostMigration-5919",
    targetNamespace:
      "/subscriptions/SubscriptionId/resourceGroups/ResourceGroup/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-4028",
  };
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const result = await client.migrationConfigs.beginCreateAndStartMigrationAndWait(
    resourceGroupName,
    namespaceName,
    configName,
    parameters
  );
  console.log(result);
}
