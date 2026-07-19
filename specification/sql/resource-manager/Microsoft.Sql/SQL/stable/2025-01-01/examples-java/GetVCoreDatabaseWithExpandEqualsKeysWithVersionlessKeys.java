
/**
 * Samples for Databases Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/GetVCoreDatabaseWithExpandEqualsKeysWithVersionlessKeys.json
     */
    /**
     * Sample code: Gets a database with database level keys expanded using versionless keys.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsADatabaseWithDatabaseLevelKeysExpandedUsingVersionlessKeys(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabases().getWithResponse("Default-SQL-SouthEastAsia", "testsvr", "testdb", "keys",
            null, com.azure.core.util.Context.NONE);
    }
}
