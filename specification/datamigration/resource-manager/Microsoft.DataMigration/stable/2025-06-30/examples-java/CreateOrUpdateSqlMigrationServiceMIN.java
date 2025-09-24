
/**
 * Samples for SqlMigrationServices CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2025-06-30/examples/
     * CreateOrUpdateSqlMigrationServiceMIN.json
     */
    /**
     * Sample code: Create or Update SQL Migration Service with minimum parameters.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void createOrUpdateSQLMigrationServiceWithMinimumParameters(
        com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.sqlMigrationServices().define("testagent").withRegion("northeurope").withExistingResourceGroup("testrg")
            .create();
    }
}
