
/**
 * Samples for WorkloadClassifiers Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/GetWorkloadClassifier.json
     */
    /**
     * Sample code: Gets a workload classifier for a data warehouse.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getsAWorkloadClassifierForADataWarehouse(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getWorkloadClassifiers().getWithResponse("Default-SQL-SouthEastAsia", "testsvr",
            "testdb", "wlm_workloadgroup", "wlm_classifier", com.azure.core.util.Context.NONE);
    }
}
