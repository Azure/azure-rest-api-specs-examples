
/**
 * Samples for ServerAdvisors Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ServerAdvisorGet.json
     */
    /**
     * Sample code: Get server advisor.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getServerAdvisor(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerAdvisors().getWithResponse("workloadinsight-demos",
            "misosisvr", "CreateIndex", com.azure.core.util.Context.NONE);
    }
}
