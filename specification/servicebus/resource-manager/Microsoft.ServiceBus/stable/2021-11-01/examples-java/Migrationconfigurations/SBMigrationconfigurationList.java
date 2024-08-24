
/**
 * Samples for MigrationConfigs List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Migrationconfigurations
     * /SBMigrationconfigurationList.json
     */
    /**
     * Sample code: MigrationConfigurationsList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void migrationConfigurationsList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getMigrationConfigs().list("ResourceGroup",
            "sdk-Namespace-9259", com.azure.core.util.Context.NONE);
    }
}
