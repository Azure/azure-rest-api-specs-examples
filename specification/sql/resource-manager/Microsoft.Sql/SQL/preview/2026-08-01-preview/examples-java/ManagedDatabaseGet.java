
/**
 * Samples for ManagedDatabases Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ManagedDatabaseGet.json
     */
    /**
     * Sample code: Gets a managed database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsAManagedDatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabases().getWithResponse("Test1", "managedInstance", "managedDatabase",
            com.azure.core.util.Context.NONE);
    }
}
