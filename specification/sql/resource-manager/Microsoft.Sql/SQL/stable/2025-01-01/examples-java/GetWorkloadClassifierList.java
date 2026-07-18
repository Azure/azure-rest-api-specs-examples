
/**
 * Samples for WorkloadClassifiers ListByWorkloadGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/GetWorkloadClassifierList.json
     */
    /**
     * Sample code: Get the list of workload classifiers for a workload group.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getTheListOfWorkloadClassifiersForAWorkloadGroup(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getWorkloadClassifiers().listByWorkloadGroup("Default-SQL-SouthEastAsia", "testsvr",
            "testdb", "wlm_workloadgroup", com.azure.core.util.Context.NONE);
    }
}
