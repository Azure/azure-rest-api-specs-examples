import com.azure.core.util.Context;

/** Samples for Databases Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-10-01-preview/examples/GetVCoreDatabase.json
     */
    /**
     * Sample code: Gets a database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsADatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getDatabases()
            .getWithResponse("Default-SQL-SouthEastAsia", "testsvr", "testdb", Context.NONE);
    }
}
