
/**
 * Samples for Servers MigrateNetworkMode.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/ServersMigrateNetworkMode.json
     */
    /**
     * Sample code: Migrate server network configuration.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void migrateServerNetworkConfiguration(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.servers().migrateNetworkMode("exampleresourcegroup", "exampleserver", com.azure.core.util.Context.NONE);
    }
}
