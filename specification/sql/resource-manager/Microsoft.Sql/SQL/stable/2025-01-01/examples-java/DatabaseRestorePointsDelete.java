
/**
 * Samples for RestorePoints Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DatabaseRestorePointsDelete.json
     */
    /**
     * Sample code: Deletes a restore point.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deletesARestorePoint(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getRestorePoints().deleteWithResponse("Default-SQL-SouthEastAsia", "testserver",
            "testDatabase", "131546477590000000", com.azure.core.util.Context.NONE);
    }
}
