
/**
 * Samples for ServerAdvisors Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerAdvisorGet.json
     */
    /**
     * Sample code: Get server advisor.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getServerAdvisor(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerAdvisors().getWithResponse("workloadinsight-demos", "misosisvr", "CreateIndex",
            com.azure.core.util.Context.NONE);
    }
}
