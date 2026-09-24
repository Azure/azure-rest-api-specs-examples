
/**
 * Samples for DatabaseSchemas ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/DatabaseSchemaListByDatabase.json
     */
    /**
     * Sample code: List database schemas.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listDatabaseSchemas(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabaseSchemas().listByDatabase("myRG", "serverName", "myDatabase", null,
            com.azure.core.util.Context.NONE);
    }
}
