
/**
 * Samples for DatabaseSchemas Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DatabaseSchemaGet.json
     */
    /**
     * Sample code: Get database schema.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getDatabaseSchema(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabaseSchemas().getWithResponse("myRG", "serverName", "myDatabase", "dbo",
            com.azure.core.util.Context.NONE);
    }
}
