
/**
 * Samples for RestorableDroppedManagedDatabases Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/GetRestorableDroppedManagedDatabase.
     * json
     */
    /**
     * Sample code: Gets a restorable dropped managed database.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsARestorableDroppedManagedDatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getRestorableDroppedManagedDatabases().getWithResponse("Test1",
            "managedInstance", "testdb,131403269876900000", com.azure.core.util.Context.NONE);
    }
}
