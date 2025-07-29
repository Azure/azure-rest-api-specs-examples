
/**
 * Samples for DatabaseMigrationsSqlDb Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/
     * SqlDbDeleteDatabaseMigration.json
     */
    /**
     * Sample code: Delete Database Migration resource.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void
        deleteDatabaseMigrationResource(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.databaseMigrationsSqlDbs().delete("testrg", "sqldbinstance", "db1", null,
            com.azure.core.util.Context.NONE);
    }
}
