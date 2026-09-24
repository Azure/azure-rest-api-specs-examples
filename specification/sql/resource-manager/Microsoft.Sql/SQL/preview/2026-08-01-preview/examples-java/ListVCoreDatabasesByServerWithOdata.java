
/**
 * Samples for Databases ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ListVCoreDatabasesByServerWithOdata.json
     */
    /**
     * Sample code: Gets a list of databases with OData filtering.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsAListOfDatabasesWithODataFiltering(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabases().listByServer("Default-SQL-SouthEastAsia", "testsvr", null, null, null,
            null, com.azure.core.util.Context.NONE);
    }
}
