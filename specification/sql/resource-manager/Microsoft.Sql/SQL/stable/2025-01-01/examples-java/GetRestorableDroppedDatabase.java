
/**
 * Samples for RestorableDroppedDatabases Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/GetRestorableDroppedDatabase.json
     */
    /**
     * Sample code: Gets a restorable dropped database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsARestorableDroppedDatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getRestorableDroppedDatabases().getWithResponse("Default-SQL-SouthEastAsia", "testsvr",
            "testdb,131403269876900000", null, null, com.azure.core.util.Context.NONE);
    }
}
