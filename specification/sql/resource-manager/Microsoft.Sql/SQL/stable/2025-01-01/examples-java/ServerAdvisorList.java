
/**
 * Samples for ServerAdvisors ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerAdvisorList.json
     */
    /**
     * Sample code: List of server advisors.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listOfServerAdvisors(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerAdvisors().listByServerWithResponse("workloadinsight-demos", "misosisvr", null,
            com.azure.core.util.Context.NONE);
    }
}
