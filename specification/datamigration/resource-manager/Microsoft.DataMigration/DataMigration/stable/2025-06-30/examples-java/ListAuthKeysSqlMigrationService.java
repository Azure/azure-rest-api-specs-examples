
/**
 * Samples for SqlMigrationServices ListAuthKeys.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/DataMigration/stable/2025-06-30/examples/
     * ListAuthKeysSqlMigrationService.json
     */
    /**
     * Sample code: Retrieve the List of Authentication Keys.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void
        retrieveTheListOfAuthenticationKeys(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.sqlMigrationServices().listAuthKeysWithResponse("testrg", "service1", com.azure.core.util.Context.NONE);
    }
}
