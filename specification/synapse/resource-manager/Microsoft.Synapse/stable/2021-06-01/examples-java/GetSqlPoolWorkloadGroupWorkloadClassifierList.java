/** Samples for SqlPoolWorkloadClassifier List. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetSqlPoolWorkloadGroupWorkloadClassifierList.json
     */
    /**
     * Sample code: Get the list of workload classifiers of a SQL Analytics pool's workload group.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getTheListOfWorkloadClassifiersOfASQLAnalyticsPoolSWorkloadGroup(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolWorkloadClassifiers()
            .list(
                "sqlcrudtest-6852",
                "sqlcrudtest-2080",
                "sqlcrudtest-9187",
                "wlm_workloadgroup",
                com.azure.core.util.Context.NONE);
    }
}
