
/**
 * Samples for WorkloadClassifiers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DeleteWorkloadClassifier.json
     */
    /**
     * Sample code: Delete a workload classifier.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deleteAWorkloadClassifier(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getWorkloadClassifiers().delete("Default-SQL-SouthEastAsia", "testsvr", "testdb",
            "wlm_workloadgroup", "wlm_workloadclassifier", com.azure.core.util.Context.NONE);
    }
}
