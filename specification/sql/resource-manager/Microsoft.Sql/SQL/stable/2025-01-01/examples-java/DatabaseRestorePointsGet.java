
/**
 * Samples for RestorePoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DatabaseRestorePointsGet.json
     */
    /**
     * Sample code: Gets a database restore point.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsADatabaseRestorePoint(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getRestorePoints().getWithResponse("Default-SQL-SouthEastAsia", "testserver",
            "testDatabase", "131546477590000000", com.azure.core.util.Context.NONE);
    }
}
