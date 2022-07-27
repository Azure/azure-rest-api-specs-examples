import com.azure.core.util.Context;

/** Samples for Databases ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-10-01-preview/examples/ListVCoreDatabasesByServer.json
     */
    /**
     * Sample code: Gets a list of databases.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsAListOfDatabases(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getDatabases()
            .listByServer("Default-SQL-SouthEastAsia", "testsvr", Context.NONE);
    }
}
