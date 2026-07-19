
/**
 * Samples for RestorableDroppedManagedDatabases Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/GetRestorableDroppedManagedDatabase.json
     */
    /**
     * Sample code: Gets a restorable dropped managed database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsARestorableDroppedManagedDatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getRestorableDroppedManagedDatabases().getWithResponse("Test1", "managedInstance",
            "testdb,131403269876900000", com.azure.core.util.Context.NONE);
    }
}
