import com.azure.core.util.Context;

/** Samples for SqlPoolWorkloadGroup Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DeleteSqlPoolWorkloadGroup.json
     */
    /**
     * Sample code: Delete a workload group of a SQL Analytics pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void deleteAWorkloadGroupOfASQLAnalyticsPool(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolWorkloadGroups()
            .delete("sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-9187", "wlm_workloadgroup", Context.NONE);
    }
}
