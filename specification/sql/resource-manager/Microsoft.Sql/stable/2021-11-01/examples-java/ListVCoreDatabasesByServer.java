
/**
 * Samples for Databases ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ListVCoreDatabasesByServer.json
     */
    /**
     * Sample code: Gets a list of databases.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsAListOfDatabases(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabases().listByServer("Default-SQL-SouthEastAsia", "testsvr",
            null, com.azure.core.util.Context.NONE);
    }
}
