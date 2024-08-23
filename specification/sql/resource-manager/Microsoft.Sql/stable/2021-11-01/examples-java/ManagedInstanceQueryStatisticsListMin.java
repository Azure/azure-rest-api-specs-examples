
import com.azure.resourcemanager.sql.models.QueryTimeGrainType;

/**
 * Samples for ManagedDatabaseQueries ListByQuery.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedInstanceQueryStatisticsListMin
     * .json
     */
    /**
     * Sample code: Obtain query execution statistics. Minimal example with only mandatory request parameters.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void obtainQueryExecutionStatisticsMinimalExampleWithOnlyMandatoryRequestParameters(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedDatabaseQueries().listByQuery("sqlcrudtest-7398",
            "sqlcrudtest-4645", "database_1", "42", null, null, QueryTimeGrainType.PT1H,
            com.azure.core.util.Context.NONE);
    }
}
