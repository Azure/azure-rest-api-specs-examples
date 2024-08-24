
import com.azure.resourcemanager.sql.models.QueryTimeGrainType;

/**
 * Samples for ManagedDatabaseQueries ListByQuery.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedInstanceQueryStatisticsListMax
     * .json
     */
    /**
     * Sample code: Obtain query execution statistics. Example with all request parameters.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void obtainQueryExecutionStatisticsExampleWithAllRequestParameters(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedDatabaseQueries().listByQuery("sqlcrudtest-7398",
            "sqlcrudtest-4645", "database_1", "42", "03/01/2020 16:23:09", "03/11/2020 14:00:00",
            QueryTimeGrainType.P1D, com.azure.core.util.Context.NONE);
    }
}
