
/**
 * Samples for Databases ListByElasticPool.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ListDatabasesByElasticPool.json
     */
    /**
     * Sample code: Gets a list of databases in an elastic pool.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsAListOfDatabasesInAnElasticPool(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabases().listByElasticPool("Default-SQL-SouthEastAsia", "testsvr", "pool1",
            com.azure.core.util.Context.NONE);
    }
}
