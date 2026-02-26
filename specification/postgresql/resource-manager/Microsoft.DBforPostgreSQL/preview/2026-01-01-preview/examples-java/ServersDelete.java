
/**
 * Samples for Servers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/ServersDelete.json
     */
    /**
     * Sample code: Delete or drop an existing server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void
        deleteOrDropAnExistingServer(com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.servers().delete("exampleresourcegroup", "exampleserver", com.azure.core.util.Context.NONE);
    }
}
