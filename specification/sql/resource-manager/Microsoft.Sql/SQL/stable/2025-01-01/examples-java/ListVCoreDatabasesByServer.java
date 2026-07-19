
/**
 * Samples for Databases ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ListVCoreDatabasesByServer.json
     */
    /**
     * Sample code: Gets a list of databases.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsAListOfDatabases(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabases().listByServer("Default-SQL-SouthEastAsia", "testsvr", null,
            com.azure.core.util.Context.NONE);
    }
}
