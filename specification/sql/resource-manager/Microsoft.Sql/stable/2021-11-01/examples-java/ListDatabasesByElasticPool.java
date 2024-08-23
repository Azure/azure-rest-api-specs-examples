
/**
 * Samples for Databases ListByElasticPool.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ListDatabasesByElasticPool.json
     */
    /**
     * Sample code: Gets a list of databases in an elastic pool.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsAListOfDatabasesInAnElasticPool(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabases().listByElasticPool("Default-SQL-SouthEastAsia",
            "testsvr", "pool1", com.azure.core.util.Context.NONE);
    }
}
