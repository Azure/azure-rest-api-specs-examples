
import com.azure.resourcemanager.sql.models.MetricType;
import com.azure.resourcemanager.sql.models.QueryTimeGrainType;

/**
 * Samples for ManagedInstances ListByManagedInstance.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceTopQueriesList.json
     */
    /**
     * Sample code: Obtain list of instance's top resource consuming queries.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        obtainListOfInstanceSTopResourceConsumingQueries(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstances().listByManagedInstance("sqlcrudtest-7398", "sqlcrudtest-4645",
            null, null, null, null, QueryTimeGrainType.PT1H, null, MetricType.DURATION,
            com.azure.core.util.Context.NONE);
    }
}
