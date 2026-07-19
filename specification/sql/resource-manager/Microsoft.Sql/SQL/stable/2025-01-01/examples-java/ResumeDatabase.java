
/**
 * Samples for Databases Resume.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ResumeDatabase.json
     */
    /**
     * Sample code: Resumes a database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void resumesADatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabases().resume("Default-SQL-SouthEastAsia", "testsvr", "testdwdb",
            com.azure.core.util.Context.NONE);
    }
}
