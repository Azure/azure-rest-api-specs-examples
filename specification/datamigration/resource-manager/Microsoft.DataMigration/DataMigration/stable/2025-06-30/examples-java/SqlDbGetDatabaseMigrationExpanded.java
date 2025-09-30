
/**
 * Samples for DatabaseMigrationsSqlDb Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/DataMigration/stable/2025-06-30/examples/
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
