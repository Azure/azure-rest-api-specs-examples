
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.MetricType;
import com.azure.resourcemanager.sql.models.QueryTimeGrainType;

/** Samples for ManagedInstances ListByManagedInstance. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedInstanceTopQueriesListMax.json
     */
    /**
     * Sample code: Obtain list of instance's top resource consuming queries. Full-blown request and response.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void obtainListOfInstanceSTopResourceConsumingQueriesFullBlownRequestAndResponse(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedInstances().listByManagedInstance("sqlcrudtest-7398",
            "sqlcrudtest-4645", null, "db1,db2", "2020-03-10T12:00:00Z", "2020-03-12T12:00:00Z", QueryTimeGrainType.P1D,
            null, MetricType.CPU, Context.NONE);
    }
}
