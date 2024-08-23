
/**
 * Samples for JobAgents Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/GetJobAgent.json
     */
    /**
     * Sample code: Get a job agent.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAJobAgent(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getJobAgents().getWithResponse("group1", "server1", "agent1",
            com.azure.core.util.Context.NONE);
    }
}
