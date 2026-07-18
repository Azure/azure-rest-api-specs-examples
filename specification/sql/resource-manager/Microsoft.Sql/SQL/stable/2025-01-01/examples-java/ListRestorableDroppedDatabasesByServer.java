
/**
 * Samples for RestorableDroppedDatabases ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ListRestorableDroppedDatabasesByServer.json
     */
    /**
     * Sample code: Gets a list of restorable dropped databases.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsAListOfRestorableDroppedDatabases(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getRestorableDroppedDatabases().listByServer("Default-SQL-SouthEastAsia", "testsvr",
            com.azure.core.util.Context.NONE);
    }
}
