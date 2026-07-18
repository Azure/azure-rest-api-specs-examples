
/**
 * Samples for PrivateLinkResources ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/PrivateLinkResourcesList.json
     */
    /**
     * Sample code: Gets private link resources for SQL.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsPrivateLinkResourcesForSQL(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getPrivateLinkResources().listByServer("Default", "test-svr",
            com.azure.core.util.Context.NONE);
    }
}
