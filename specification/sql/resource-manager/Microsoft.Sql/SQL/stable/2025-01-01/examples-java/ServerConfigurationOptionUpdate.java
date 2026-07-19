
import com.azure.resourcemanager.sql.fluent.models.ServerConfigurationOptionInner;
import com.azure.resourcemanager.sql.models.ServerConfigurationOptionName;

/**
 * Samples for ServerConfigurationOptions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerConfigurationOptionUpdate.json
     */
    /**
     * Sample code: Updates managed instance server configuration option.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        updatesManagedInstanceServerConfigurationOption(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerConfigurationOptions().createOrUpdate("testrg", "testinstance",
            ServerConfigurationOptionName.ALLOW_POLYBASE_EXPORT,
            new ServerConfigurationOptionInner().withServerConfigurationOptionValue(1),
            com.azure.core.util.Context.NONE);
    }
}
