
/**
 * Samples for Servers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ServerDelete.json
     */
    /**
     * Sample code: Delete server.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deleteServer(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServers().delete("sqlcrudtest-7398", "sqlcrudtest-6661",
            com.azure.core.util.Context.NONE);
    }
}
