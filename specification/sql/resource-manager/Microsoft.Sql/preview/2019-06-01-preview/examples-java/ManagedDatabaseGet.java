import com.azure.core.util.Context;

/** Samples for ManagedDatabases Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2019-06-01-preview/examples/ManagedDatabaseGet.json
     */
    /**
     * Sample code: Gets a managed database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsAManagedDatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getManagedDatabases()
            .getWithResponse("Test1", "managedInstance", "managedDatabase", Context.NONE);
    }
}
