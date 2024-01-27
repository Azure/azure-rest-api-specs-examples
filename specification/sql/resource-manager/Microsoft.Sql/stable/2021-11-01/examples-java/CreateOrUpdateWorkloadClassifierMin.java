
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.WorkloadClassifierInner;

/** Samples for WorkloadClassifiers CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/CreateOrUpdateWorkloadClassifierMin.
     * json
     */
    /**
     * Sample code: Create a workload group with the required properties specified.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createAWorkloadGroupWithTheRequiredPropertiesSpecified(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getWorkloadClassifiers().createOrUpdate(
            "Default-SQL-SouthEastAsia", "testsvr", "testdb", "wlm_workloadgroup", "wlm_workloadclassifier",
            new WorkloadClassifierInner().withMemberName("dbo"), Context.NONE);
    }
}
