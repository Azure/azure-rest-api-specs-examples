
import com.azure.resourcemanager.servicebus.models.MigrationConfigurationName;

/**
 * Samples for MigrationConfigs Revert.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Migrationconfigurations/SBMigrationconfigurationRevert.json
     */
    /**
     * Sample code: MigrationConfigurationsRevert.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void migrationConfigurationsRevert(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getMigrationConfigs().revertWithResponse("ResourceGroup", "sdk-Namespace-41",
            MigrationConfigurationName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
