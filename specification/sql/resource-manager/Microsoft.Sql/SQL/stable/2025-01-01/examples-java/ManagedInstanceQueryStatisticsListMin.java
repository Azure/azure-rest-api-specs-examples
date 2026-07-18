
import com.azure.resourcemanager.sql.models.QueryTimeGrainType;

/**
 * Samples for ManagedDatabaseQueries ListByQuery.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceQueryStatisticsListMin.json
     */
    /**
     * Sample code: Obtain query execution statistics. Minimal example with only mandatory request parameters.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void obtainQueryExecutionStatisticsMinimalExampleWithOnlyMandatoryRequestParameters(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabaseQueries().listByQuery("sqlcrudtest-7398", "sqlcrudtest-4645",
            "database_1", "42", null, null, QueryTimeGrainType.PT1H, com.azure.core.util.Context.NONE);
    }
}
