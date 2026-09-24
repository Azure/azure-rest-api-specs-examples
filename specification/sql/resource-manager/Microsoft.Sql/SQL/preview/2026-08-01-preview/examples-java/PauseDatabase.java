
/**
 * Samples for Databases Pause.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/PauseDatabase.json
     */
    /**
     * Sample code: Pauses a database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void pausesADatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabases().pause("Default-SQL-SouthEastAsia", "testsvr", "testdwdb",
            com.azure.core.util.Context.NONE);
    }
}
