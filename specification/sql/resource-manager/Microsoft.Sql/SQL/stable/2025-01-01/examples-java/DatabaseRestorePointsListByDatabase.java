
/**
 * Samples for RestorePoints ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DatabaseRestorePointsListByDatabase.json
     */
    /**
     * Sample code: List database restore points.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listDatabaseRestorePoints(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getRestorePoints().listByDatabase("sqlcrudtest-6730", "sqlcrudtest-9007", "3481",
            com.azure.core.util.Context.NONE);
    }
}
