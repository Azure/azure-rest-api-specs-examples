
import com.azure.core.util.Context;
import com.azure.resourcemanager.servicebus.models.MigrationConfigurationName;

/** Samples for MigrationConfigs CompleteMigration. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Migrationconfigurations
     * /SBMigrationconfigurationCompleteMigration.json
     */
    /**
     * Sample code: MigrationConfigurationsCompleteMigration.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void migrationConfigurationsCompleteMigration(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getMigrationConfigs().completeMigrationWithResponse(
            "ResourceGroup", "sdk-Namespace-41", MigrationConfigurationName.DEFAULT, Context.NONE);
    }
}
