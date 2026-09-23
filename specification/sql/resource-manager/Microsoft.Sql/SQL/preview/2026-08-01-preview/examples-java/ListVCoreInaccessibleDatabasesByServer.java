
/**
 * Samples for Databases ListInaccessibleByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ListVCoreInaccessibleDatabasesByServer.json
     */
    /**
     * Sample code: Gets a list of inaccessible databases in a logical server.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getsAListOfInaccessibleDatabasesInALogicalServer(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabases().listInaccessibleByServer("Default-SQL-SouthEastAsia", "testsvr",
            com.azure.core.util.Context.NONE);
    }
}
