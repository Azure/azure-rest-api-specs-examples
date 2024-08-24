
/**
 * Samples for DatabaseUsages ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/GetDatabaseUsages.json
     */
    /**
     * Sample code: Gets database usages.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsDatabaseUsages(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabaseUsages().listByDatabase("Default-SQL-SouthEastAsia",
            "testsvr", "testdb", com.azure.core.util.Context.NONE);
    }
}
