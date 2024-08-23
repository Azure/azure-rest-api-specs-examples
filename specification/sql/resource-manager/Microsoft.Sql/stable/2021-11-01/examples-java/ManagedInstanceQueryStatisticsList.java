
/**
 * Samples for ManagedDatabaseQueries ListByQuery.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedInstanceQueryStatisticsList.
     * json
     */
    /**
     * Sample code: Obtain query execution statistics.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void obtainQueryExecutionStatistics(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedDatabaseQueries().listByQuery("sqlcrudtest-7398",
            "sqlcrudtest-4645", "database_1", "42", null, null, null, com.azure.core.util.Context.NONE);
    }
}
