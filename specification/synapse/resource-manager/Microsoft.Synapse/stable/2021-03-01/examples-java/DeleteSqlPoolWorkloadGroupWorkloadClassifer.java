
/**
 * Samples for SqlPoolWorkloadClassifier Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-03-01/examples/
     * DeleteSqlPoolWorkloadGroupWorkloadClassifer.json
     */
    /**
     * Sample code: Delete a workload classifier of a SQL Analytics pool's workload group.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void deleteAWorkloadClassifierOfASQLAnalyticsPoolSWorkloadGroup(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.sqlPoolWorkloadClassifiers().delete("sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-9187",
            "wlm_workloadgroup", "wlm_workloadclassifier", com.azure.core.util.Context.NONE);
    }
}
