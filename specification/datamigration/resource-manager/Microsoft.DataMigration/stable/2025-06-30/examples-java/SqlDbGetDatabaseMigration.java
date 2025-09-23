
/**
 * Samples for DatabaseMigrationsSqlDb Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2025-06-30/examples/
     * SqlDbGetDatabaseMigration.json
     */
    /**
     * Sample code: Get Sql DB database Migration without the expand parameter.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void getSqlDBDatabaseMigrationWithoutTheExpandParameter(
        com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.databaseMigrationsSqlDbs().getWithResponse("testrg", "sqldbinstance", "db1", null, null,
            com.azure.core.util.Context.NONE);
    }
}
