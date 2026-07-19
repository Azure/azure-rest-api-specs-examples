
/**
 * Samples for DatabaseTables Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DatabaseTableGet.json
     */
    /**
     * Sample code: Get database table.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getDatabaseTable(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabaseTables().getWithResponse("myRG", "serverName", "myDatabase", "dbo", "table1",
            com.azure.core.util.Context.NONE);
    }
}
