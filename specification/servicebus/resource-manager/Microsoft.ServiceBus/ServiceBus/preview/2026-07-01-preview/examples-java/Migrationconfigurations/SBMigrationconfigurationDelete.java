
import com.azure.resourcemanager.servicebus.models.MigrationConfigurationName;

/**
 * Samples for MigrationConfigs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Migrationconfigurations/SBMigrationconfigurationDelete.json
     */
    /**
     * Sample code: MigrationConfigurationsDelete.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void migrationConfigurationsDelete(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getMigrationConfigs().deleteWithResponse("ResourceGroup", "sdk-Namespace-41",
            MigrationConfigurationName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
