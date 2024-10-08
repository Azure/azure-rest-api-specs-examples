
/**
 * Samples for DatabaseSchemas Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DatabaseSchemaGet.json
     */
    /**
     * Sample code: Get database schema.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDatabaseSchema(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabaseSchemas().getWithResponse("myRG", "serverName",
            "myDatabase", "dbo", com.azure.core.util.Context.NONE);
    }
}
