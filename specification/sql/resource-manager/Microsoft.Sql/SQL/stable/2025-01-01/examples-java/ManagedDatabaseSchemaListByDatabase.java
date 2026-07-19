
/**
 * Samples for ManagedDatabaseSchemas ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedDatabaseSchemaListByDatabase.json
     */
    /**
     * Sample code: List managed database schemas.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listManagedDatabaseSchemas(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabaseSchemas().listByDatabase("myRG", "myManagedInstanceName",
            "myDatabase", null, com.azure.core.util.Context.NONE);
    }
}
