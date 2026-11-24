
/**
 * Samples for MigrationConfigs List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/Migrationconfigurations/SBMigrationconfigurationList.json
     */
    /**
     * Sample code: MigrationConfigurationsList.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void migrationConfigurationsList(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getMigrationConfigs().list("ResourceGroup", "sdk-Namespace-9259",
            com.azure.core.util.Context.NONE);
    }
}
