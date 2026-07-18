
/**
 * Samples for ManagedDatabases ListInaccessibleByInstance.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/InaccessibleManagedDatabaseListByManagedInstance.json
     */
    /**
     * Sample code: List inaccessible managed databases by managed instances.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        listInaccessibleManagedDatabasesByManagedInstances(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabases().listInaccessibleByInstance("testrg", "testcl",
            com.azure.core.util.Context.NONE);
    }
}
