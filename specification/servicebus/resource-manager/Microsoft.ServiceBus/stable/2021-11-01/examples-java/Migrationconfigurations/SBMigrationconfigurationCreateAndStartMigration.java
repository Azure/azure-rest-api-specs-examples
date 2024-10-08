
import com.azure.resourcemanager.servicebus.fluent.models.MigrationConfigPropertiesInner;
import com.azure.resourcemanager.servicebus.models.MigrationConfigurationName;

/**
 * Samples for MigrationConfigs CreateAndStartMigration.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Migrationconfigurations
     * /SBMigrationconfigurationCreateAndStartMigration.json
     */
    /**
     * Sample code: MigrationConfigurationsStartMigration.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void migrationConfigurationsStartMigration(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getMigrationConfigs().createAndStartMigration(
            "ResourceGroup", "sdk-Namespace-41", MigrationConfigurationName.DEFAULT,
            new MigrationConfigPropertiesInner().withTargetNamespace(
                "/subscriptions/SubscriptionId/resourceGroups/ResourceGroup/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-4028")
                .withPostMigrationName("sdk-PostMigration-5919"),
            com.azure.core.util.Context.NONE);
    }
}
