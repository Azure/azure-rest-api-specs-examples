
/**
 * Samples for GeoBackupPolicies List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/GeoBackupPoliciesList.json
     */
    /**
     * Sample code: List Geo backup policies for the given database resource.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        listGeoBackupPoliciesForTheGivenDatabaseResource(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getGeoBackupPolicies().list("sqlcrudtest-4799", "sqlcrudtest-5961", "testdw",
            com.azure.core.util.Context.NONE);
    }
}
