
/**
 * Samples for ServerAdvisors ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ServerRecommendedActionListExpand.
     * json
     */
    /**
     * Sample code: List of server recommended actions for all advisors.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listOfServerRecommendedActionsForAllAdvisors(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerAdvisors().listByServerWithResponse(
            "workloadinsight-demos", "misosisvr", "recommendedActions", com.azure.core.util.Context.NONE);
    }
}
