const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a MigrationConfiguration
 *
 * @summary Deletes a MigrationConfiguration
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Migrationconfigurations/SBMigrationconfigurationDelete.json
 */
async function migrationConfigurationsDelete() {
  const subscriptionId = "SubscriptionId";
  const resourceGroupName = "ResourceGroup";
  const namespaceName = "sdk-Namespace-41";
  const configName = "$default";
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const result = await client.migrationConfigs.delete(resourceGroupName, namespaceName, configName);
  console.log(result);
}

migrationConfigurationsDelete().catch(console.error);
