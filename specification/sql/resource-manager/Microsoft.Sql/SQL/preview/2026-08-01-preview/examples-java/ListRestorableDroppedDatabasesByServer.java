
/**
 * Samples for RestorableDroppedDatabases ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ListRestorableDroppedDatabasesByServer.json
     */
    /**
     * Sample code: Gets a list of restorable dropped databases.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsAListOfRestorableDroppedDatabases(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getRestorableDroppedDatabases().listByServer("Default-SQL-SouthEastAsia", "testsvr",
            null, null, com.azure.core.util.Context.NONE);
    }
}
