
/**
 * Samples for Servers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerList.json
     */
    /**
     * Sample code: List servers.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listServers(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServers().list(null, com.azure.core.util.Context.NONE);
    }
}
