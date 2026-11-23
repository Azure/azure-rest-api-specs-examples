
import com.azure.resourcemanager.servicebus.models.MigrationConfigurationName;

/**
 * Samples for MigrationConfigs CompleteMigration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Migrationconfigurations/SBMigrationconfigurationCompleteMigration.json
     */
    /**
     * Sample code: MigrationConfigurationsCompleteMigration.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void
        migrationConfigurationsCompleteMigration(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getMigrationConfigs().completeMigrationWithResponse("ResourceGroup", "sdk-Namespace-41",
            MigrationConfigurationName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
