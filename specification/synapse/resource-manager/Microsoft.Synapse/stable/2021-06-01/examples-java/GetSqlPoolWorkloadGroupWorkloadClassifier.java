import com.azure.core.util.Context;

/** Samples for SqlPoolWorkloadClassifier Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetSqlPoolWorkloadGroupWorkloadClassifier.json
     */
    /**
     * Sample code: Get a workload classifier for SQL Analytics pool's workload group.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getAWorkloadClassifierForSQLAnalyticsPoolSWorkloadGroup(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolWorkloadClassifiers()
            .getWithResponse(
                "sqlcrudtest-6852",
                "sqlcrudtest-2080",
                "sqlcrudtest-9187",
                "wlm_workloadgroup",
                "wlm_classifier",
                Context.NONE);
    }
}
