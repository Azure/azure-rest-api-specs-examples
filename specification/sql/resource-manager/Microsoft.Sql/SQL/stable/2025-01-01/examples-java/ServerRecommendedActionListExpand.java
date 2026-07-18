
/**
 * Samples for ServerAdvisors ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerRecommendedActionListExpand.json
     */
    /**
     * Sample code: List of server recommended actions for all advisors.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        listOfServerRecommendedActionsForAllAdvisors(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerAdvisors().listByServerWithResponse("workloadinsight-demos", "misosisvr",
            "recommendedActions", com.azure.core.util.Context.NONE);
    }
}
