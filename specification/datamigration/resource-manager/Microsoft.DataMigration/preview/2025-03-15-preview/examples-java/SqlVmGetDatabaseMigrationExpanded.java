
/**
 * Samples for DatabaseMigrationsSqlVm Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/
     * SqlVmGetDatabaseMigrationExpanded.json
     */
    /**
     * Sample code: Get Sql VM database Migration with the expand parameter.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void getSqlVMDatabaseMigrationWithTheExpandParameter(
        com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.databaseMigrationsSqlVms().getWithResponse("testrg", "testvm", "db1", null, "MigrationStatusDetails",
            com.azure.core.util.Context.NONE);
    }
}
