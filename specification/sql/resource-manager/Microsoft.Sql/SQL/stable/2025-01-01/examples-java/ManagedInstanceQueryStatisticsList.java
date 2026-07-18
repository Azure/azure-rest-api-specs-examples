
/**
 * Samples for ManagedDatabaseQueries ListByQuery.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceQueryStatisticsList.json
     */
    /**
     * Sample code: Obtain query execution statistics.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void obtainQueryExecutionStatistics(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabaseQueries().listByQuery("sqlcrudtest-7398", "sqlcrudtest-4645",
            "database_1", "42", null, null, null, com.azure.core.util.Context.NONE);
    }
}
