
import com.azure.core.util.Context;

/** Samples for WorkloadClassifiers ListByWorkloadGroup. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/GetWorkloadClassifierList.json
     */
    /**
     * Sample code: Get the list of workload classifiers for a workload group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getTheListOfWorkloadClassifiersForAWorkloadGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getWorkloadClassifiers()
            .listByWorkloadGroup("Default-SQL-SouthEastAsia", "testsvr", "testdb", "wlm_workloadgroup", Context.NONE);
    }
}
