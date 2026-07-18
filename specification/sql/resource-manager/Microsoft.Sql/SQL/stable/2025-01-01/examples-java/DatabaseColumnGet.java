
/**
 * Samples for DatabaseColumns Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DatabaseColumnGet.json
     */
    /**
     * Sample code: Get database column.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getDatabaseColumn(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabaseColumns().getWithResponse("myRG", "serverName", "myDatabase", "dbo",
            "table1", "column1", com.azure.core.util.Context.NONE);
    }
}
