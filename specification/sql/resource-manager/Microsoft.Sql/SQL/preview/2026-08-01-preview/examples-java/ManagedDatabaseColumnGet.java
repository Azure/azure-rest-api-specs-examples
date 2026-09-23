
/**
 * Samples for ManagedDatabaseColumns Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ManagedDatabaseColumnGet.json
     */
    /**
     * Sample code: Get managed database column.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getManagedDatabaseColumn(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabaseColumns().getWithResponse("myRG", "myManagedInstanceName",
            "myDatabase", "dbo", "table1", "column1", com.azure.core.util.Context.NONE);
    }
}
