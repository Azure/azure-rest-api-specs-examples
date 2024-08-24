
/**
 * Samples for RecoverableManagedDatabases Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/GetRecoverableManagedDatabase.json
     */
    /**
     * Sample code: Gets a recoverable databases by managed instances.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getsARecoverableDatabasesByManagedInstances(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getRecoverableManagedDatabases().getWithResponse("Test1",
            "managedInstance", "testdb", com.azure.core.util.Context.NONE);
    }
}
