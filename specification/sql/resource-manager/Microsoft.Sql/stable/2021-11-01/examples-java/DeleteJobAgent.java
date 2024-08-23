
/**
 * Samples for JobAgents Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DeleteJobAgent.json
     */
    /**
     * Sample code: Delete a job agent.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAJobAgent(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getJobAgents().delete("group1", "server1", "agent1",
            com.azure.core.util.Context.NONE);
    }
}
