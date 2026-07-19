
/**
 * Samples for DatabaseSchemas ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DatabaseSchemaListByDatabase.json
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
