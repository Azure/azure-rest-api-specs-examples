
/**
 * Samples for DeletedServers ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/DeletedServerList.json
     */
    /**
     * Sample code: List deleted servers.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listDeletedServers(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDeletedServers().listByLocation("japaneast", com.azure.core.util.Context.NONE);
    }
}
