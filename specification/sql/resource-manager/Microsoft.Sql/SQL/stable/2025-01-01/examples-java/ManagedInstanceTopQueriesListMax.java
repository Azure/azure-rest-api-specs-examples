
import com.azure.resourcemanager.sql.models.MetricType;
import com.azure.resourcemanager.sql.models.QueryTimeGrainType;

/**
 * Samples for ManagedInstances ListByManagedInstance.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceTopQueriesListMax.json
     */
    /**
     * Sample code: Obtain list of instance's top resource consuming queries. Full-blown request and response.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void obtainListOfInstanceSTopResourceConsumingQueriesFullBlownRequestAndResponse(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstances().listByManagedInstance("sqlcrudtest-7398", "sqlcrudtest-4645",
            null, "db1,db2", "2020-03-10T12:00:00Z", "2020-03-12T12:00:00Z", QueryTimeGrainType.P1D, null,
            MetricType.CPU, com.azure.core.util.Context.NONE);
    }
}
