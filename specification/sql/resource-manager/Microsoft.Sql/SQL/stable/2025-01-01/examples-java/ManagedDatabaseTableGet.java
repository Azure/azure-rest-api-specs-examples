
/**
 * Samples for ManagedDatabaseTables Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedDatabaseTableGet.json
     */
    /**
     * Sample code: Get managed database table.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getManagedDatabaseTable(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabaseTables().getWithResponse("myRG", "myManagedInstanceName",
            "myDatabase", "dbo", "table1", com.azure.core.util.Context.NONE);
    }
}
