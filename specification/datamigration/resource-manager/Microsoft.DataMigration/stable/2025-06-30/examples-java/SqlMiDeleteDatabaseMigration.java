
/**
 * Samples for DatabaseMigrationsSqlMi Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2025-06-30/examples/
     * SqlMiDeleteDatabaseMigration.json
     */
    /**
     * Sample code: Delete Database Migration resource with Minimum parameters.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void deleteDatabaseMigrationResourceWithMinimumParameters(
        com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.databaseMigrationsSqlMis().delete("testrg", "managedInstance1", "db1", null,
            com.azure.core.util.Context.NONE);
    }
}
