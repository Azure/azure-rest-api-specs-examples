
/**
 * Samples for ManagedDatabaseColumns ListByTable.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedDatabaseColumnListByTable.json
     */
    /**
     * Sample code: List managed database columns.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listManagedDatabaseColumns(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabaseColumns().listByTable("myRG", "myManagedInstanceName", "myDatabase",
            "dbo", "table1", null, com.azure.core.util.Context.NONE);
    }
}
