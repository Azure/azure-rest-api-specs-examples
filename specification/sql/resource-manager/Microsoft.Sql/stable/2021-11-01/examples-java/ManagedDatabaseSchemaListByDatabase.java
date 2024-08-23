
/**
 * Samples for ManagedDatabaseSchemas ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedDatabaseSchemaListByDatabase.
     * json
     */
    /**
     * Sample code: List managed database schemas.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listManagedDatabaseSchemas(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedDatabaseSchemas().listByDatabase("myRG",
            "myManagedInstanceName", "myDatabase", null, com.azure.core.util.Context.NONE);
    }
}
