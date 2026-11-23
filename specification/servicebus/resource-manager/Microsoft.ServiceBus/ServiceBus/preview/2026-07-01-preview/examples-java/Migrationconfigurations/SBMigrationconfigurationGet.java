
import com.azure.resourcemanager.servicebus.models.MigrationConfigurationName;

/**
 * Samples for MigrationConfigs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Migrationconfigurations/SBMigrationconfigurationGet.json
     */
    /**
     * Sample code: MigrationConfigurationsGet.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void migrationConfigurationsGet(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getMigrationConfigs().getWithResponse("ResourceGroup", "sdk-Namespace-41",
            MigrationConfigurationName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
