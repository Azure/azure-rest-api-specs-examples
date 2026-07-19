
/**
 * Samples for ManagedDatabaseQueries Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceQueryGet.json
     */
    /**
     * Sample code: Obtain query properties.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void obtainQueryProperties(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabaseQueries().getWithResponse("sqlcrudtest-7398", "sqlcrudtest-4645",
            "database_1", "42", com.azure.core.util.Context.NONE);
    }
}
