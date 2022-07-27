import com.azure.core.util.Context;

/** Samples for ManagedDatabases ListInaccessibleByInstance. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2019-06-01-preview/examples/InaccessibleManagedDatabaseListByManagedInstance.json
     */
    /**
     * Sample code: List inaccessible managed databases by managed instances.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listInaccessibleManagedDatabasesByManagedInstances(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getManagedDatabases()
            .listInaccessibleByInstance("testrg", "testcl", Context.NONE);
    }
}
