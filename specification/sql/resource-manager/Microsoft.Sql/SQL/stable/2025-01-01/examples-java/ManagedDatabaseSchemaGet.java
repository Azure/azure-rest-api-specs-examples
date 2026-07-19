
/**
 * Samples for ManagedDatabaseSchemas Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedDatabaseSchemaGet.json
     */
    /**
     * Sample code: Get managed database schema.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getManagedDatabaseSchema(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabaseSchemas().getWithResponse("myRG", "myManagedInstanceName",
            "myDatabase", "dbo", com.azure.core.util.Context.NONE);
    }
}
