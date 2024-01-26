
import com.azure.core.util.Context;
import com.azure.resourcemanager.servicebus.models.MigrationConfigurationName;

/** Samples for MigrationConfigs Revert. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Migrationconfigurations
     * /SBMigrationconfigurationRevert.json
     */
    /**
     * Sample code: MigrationConfigurationsRevert.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void migrationConfigurationsRevert(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getMigrationConfigs().revertWithResponse("ResourceGroup",
            "sdk-Namespace-41", MigrationConfigurationName.DEFAULT, Context.NONE);
    }
}
