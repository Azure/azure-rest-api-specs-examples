
/**
 * Samples for Servers GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerGet.json
     */
    /**
     * Sample code: Get server.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getServer(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServers().getByResourceGroupWithResponse("sqlcrudtest-7398", "sqlcrudtest-4645",
            null, com.azure.core.util.Context.NONE);
    }
}
