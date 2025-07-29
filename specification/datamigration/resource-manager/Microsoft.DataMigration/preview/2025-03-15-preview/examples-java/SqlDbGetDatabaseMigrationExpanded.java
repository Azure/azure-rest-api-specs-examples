
/**
 * Samples for DatabaseMigrationsSqlDb Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/
     * SqlDbGetDatabaseMigrationExpanded.json
     */
    /**
     * Sample code: Get Sql DB database Migration with the expand parameter.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void getSqlDBDatabaseMigrationWithTheExpandParameter(
        com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.databaseMigrationsSqlDbs().getWithResponse("testrg", "sqldbinstance", "db1", null,
            "MigrationStatusDetails", com.azure.core.util.Context.NONE);
    }
}
