
/**
 * Samples for SqlMigrationServices GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/
     * GetSqlMigrationService.json
     */
    /**
     * Sample code: Get Migration Service.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void getMigrationService(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.sqlMigrationServices().getByResourceGroupWithResponse("testrg", "service1",
            com.azure.core.util.Context.NONE);
    }
}
