
/**
 * Samples for DatabaseMigrationsSqlMi Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2025-06-30/examples/
     * SqlMiGetDatabaseMigration.json
     */
    /**
     * Sample code: Get Sql MI database Migration without the expand parameter.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void getSqlMIDatabaseMigrationWithoutTheExpandParameter(
        com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.databaseMigrationsSqlMis().getWithResponse("testrg", "managedInstance1", "db1", null, null,
            com.azure.core.util.Context.NONE);
    }
}
