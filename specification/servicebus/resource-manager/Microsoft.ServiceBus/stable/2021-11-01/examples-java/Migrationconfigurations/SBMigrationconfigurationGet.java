
import com.azure.resourcemanager.servicebus.models.MigrationConfigurationName;

/**
 * Samples for MigrationConfigs Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Migrationconfigurations
     * /SBMigrationconfigurationGet.json
     */
    /**
     * Sample code: MigrationConfigurationsGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void migrationConfigurationsGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getMigrationConfigs().getWithResponse("ResourceGroup",
            "sdk-Namespace-41", MigrationConfigurationName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
