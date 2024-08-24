
/**
 * Samples for ManagedDatabaseSchemas Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedDatabaseSchemaGet.json
     */
    /**
     * Sample code: Get managed database schema.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getManagedDatabaseSchema(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedDatabaseSchemas().getWithResponse("myRG",
            "myManagedInstanceName", "myDatabase", "dbo", com.azure.core.util.Context.NONE);
    }
}
