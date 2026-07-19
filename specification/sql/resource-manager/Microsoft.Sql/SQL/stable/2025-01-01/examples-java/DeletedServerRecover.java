
/**
 * Samples for DeletedServers Recover.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DeletedServerRecover.json
     */
    /**
     * Sample code: Recover deleted server.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void recoverDeletedServer(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDeletedServers().recover("japaneast", "sqlcrudtest-d-1414",
            com.azure.core.util.Context.NONE);
    }
}
