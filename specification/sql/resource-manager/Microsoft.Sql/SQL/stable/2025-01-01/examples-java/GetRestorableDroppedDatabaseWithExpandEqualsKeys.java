
/**
 * Samples for RestorableDroppedDatabases Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/GetRestorableDroppedDatabaseWithExpandEqualsKeys.json
     */
    /**
     * Sample code: Gets a restorable dropped database with expand equals keys.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getsARestorableDroppedDatabaseWithExpandEqualsKeys(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getRestorableDroppedDatabases().getWithResponse("Default-SQL-SouthEastAsia", "testsvr",
            "testdb,131403269876900000", "keys", null, com.azure.core.util.Context.NONE);
    }
}
