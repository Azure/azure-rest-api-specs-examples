
/**
 * Samples for DeletedServers Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DeletedServerGet.json
     */
    /**
     * Sample code: Get deleted server.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getDeletedServer(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDeletedServers().getWithResponse("japaneast", "sqlcrudtest-d-1414",
            com.azure.core.util.Context.NONE);
    }
}
