
/**
 * Samples for DatabaseAdvisors Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DatabaseAdvisorGet.json
     */
    /**
     * Sample code: Get database advisor.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getDatabaseAdvisor(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabaseAdvisors().getWithResponse("workloadinsight-demos", "misosisvr",
            "IndexAdvisor_test_3", "CreateIndex", com.azure.core.util.Context.NONE);
    }
}
