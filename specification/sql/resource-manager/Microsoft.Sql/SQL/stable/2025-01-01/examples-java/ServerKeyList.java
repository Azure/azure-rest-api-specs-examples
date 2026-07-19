
/**
 * Samples for ServerKeys ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerKeyList.json
     */
    /**
     * Sample code: List the server keys by server.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listTheServerKeysByServer(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerKeys().listByServer("sqlcrudtest-7398", "sqlcrudtest-4645",
            com.azure.core.util.Context.NONE);
    }
}
