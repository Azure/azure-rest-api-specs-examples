
/**
 * Samples for RecoverableManagedDatabases Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/GetRecoverableManagedDatabase.json
     */
    /**
     * Sample code: Gets a recoverable databases by managed instances.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getsARecoverableDatabasesByManagedInstances(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getRecoverableManagedDatabases().getWithResponse("Test1", "managedInstance", "testdb",
            com.azure.core.util.Context.NONE);
    }
}
