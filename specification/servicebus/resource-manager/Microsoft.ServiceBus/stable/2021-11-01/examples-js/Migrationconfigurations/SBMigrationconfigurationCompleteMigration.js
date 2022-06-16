const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This operation Completes Migration of entities by pointing the connection strings to Premium namespace and any entities created after the operation will be under Premium Namespace. CompleteMigration operation will fail when entity migration is in-progress.
 *
 * @summary This operation Completes Migration of entities by pointing the connection strings to Premium namespace and any entities created after the operation will be under Premium Namespace. CompleteMigration operation will fail when entity migration is in-progress.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Migrationconfigurations/SBMigrationconfigurationCompleteMigration.json
 */
async function migrationConfigurationsCompleteMigration() {
  const subscriptionId = "SubscriptionId";
  const resourceGroupName = "ResourceGroup";
  const namespaceName = "sdk-Namespace-41";
  const configName = "$default";
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const result = await client.migrationConfigs.completeMigration(
    resourceGroupName,
    namespaceName,
    configName
  );
  console.log(result);
}

migrationConfigurationsCompleteMigration().catch(console.error);
