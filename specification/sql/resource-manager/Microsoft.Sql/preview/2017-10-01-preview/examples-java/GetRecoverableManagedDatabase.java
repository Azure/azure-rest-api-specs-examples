import com.azure.core.util.Context;

/** Samples for RecoverableManagedDatabases Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-10-01-preview/examples/GetRecoverableManagedDatabase.json
     */
    /**
     * Sample code: Gets a recoverable databases by managed instances.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsARecoverableDatabasesByManagedInstances(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getRecoverableManagedDatabases()
            .getWithResponse("Test1", "managedInstance", "testdb", Context.NONE);
    }
}
