
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.MetricType;
import com.azure.resourcemanager.sql.models.QueryTimeGrainType;

/** Samples for ManagedInstances ListByManagedInstance. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedInstanceTopQueriesList.json
     */
    /**
     * Sample code: Obtain list of instance's top resource consuming queries.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        obtainListOfInstanceSTopResourceConsumingQueries(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedInstances().listByManagedInstance("sqlcrudtest-7398",
            "sqlcrudtest-4645", null, null, null, null, QueryTimeGrainType.PT1H, null, MetricType.DURATION,
            Context.NONE);
    }
}
