
/**
 * Samples for DatabaseSchemas ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DatabaseSchemaListByDatabase.json
     */
    /**
     * Sample code: List database schemas.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDatabaseSchemas(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabaseSchemas().listByDatabase("myRG", "serverName",
            "myDatabase", null, com.azure.core.util.Context.NONE);
    }
}
