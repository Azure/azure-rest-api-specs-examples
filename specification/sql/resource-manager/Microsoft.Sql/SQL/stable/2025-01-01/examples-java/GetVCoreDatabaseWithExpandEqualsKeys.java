
/**
 * Samples for Databases Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/GetVCoreDatabaseWithExpandEqualsKeys.json
     */
    /**
     * Sample code: Gets a database with database level keys expanded.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getsADatabaseWithDatabaseLevelKeysExpanded(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabases().getWithResponse("Default-SQL-SouthEastAsia", "testsvr", "testdb", "keys",
            null, com.azure.core.util.Context.NONE);
    }
}
