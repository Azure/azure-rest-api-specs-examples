
import com.azure.core.util.Context;

/** Samples for WorkloadGroups Delete. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DeleteWorkloadGroup.json
     */
    /**
     * Sample code: Delete a workload group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAWorkloadGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getWorkloadGroups().delete("Default-SQL-SouthEastAsia", "testsvr",
            "testdb", "wlm_workloadgroup", Context.NONE);
    }
}
