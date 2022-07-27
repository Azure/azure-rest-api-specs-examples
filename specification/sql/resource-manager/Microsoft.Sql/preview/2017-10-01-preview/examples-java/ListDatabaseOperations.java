import com.azure.core.util.Context;

/** Samples for DatabaseOperations ListByDatabase. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-10-01-preview/examples/ListDatabaseOperations.json
     */
    /**
     * Sample code: List the database management operations.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listTheDatabaseManagementOperations(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getDatabaseOperations()
            .listByDatabase("sqlcrudtest-7398", "sqlcrudtest-4645", "testdb", Context.NONE);
    }
}
