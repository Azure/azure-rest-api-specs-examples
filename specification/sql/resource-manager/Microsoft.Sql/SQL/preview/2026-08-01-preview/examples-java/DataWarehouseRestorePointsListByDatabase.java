
/**
 * Samples for RestorePoints ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/DataWarehouseRestorePointsListByDatabase.json
     */
    /**
     * Sample code: List datawarehouse database restore points.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listDatawarehouseDatabaseRestorePoints(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getRestorePoints().listByDatabase("Default-SQL-SouthEastAsia", "testserver",
            "testDatabase", com.azure.core.util.Context.NONE);
    }
}
