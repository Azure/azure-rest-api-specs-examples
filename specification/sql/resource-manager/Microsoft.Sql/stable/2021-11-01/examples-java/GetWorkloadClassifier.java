
import com.azure.core.util.Context;

/** Samples for WorkloadClassifiers Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/GetWorkloadClassifier.json
     */
    /**
     * Sample code: Gets a workload classifier for a data warehouse.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsAWorkloadClassifierForADataWarehouse(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getWorkloadClassifiers().getWithResponse(
            "Default-SQL-SouthEastAsia", "testsvr", "testdb", "wlm_workloadgroup", "wlm_classifier", Context.NONE);
    }
}
