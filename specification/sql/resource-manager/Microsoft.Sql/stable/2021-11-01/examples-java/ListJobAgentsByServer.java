
/**
 * Samples for JobAgents ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ListJobAgentsByServer.json
     */
    /**
     * Sample code: List job agents in a server.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listJobAgentsInAServer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getJobAgents().listByServer("group1", "server1",
            com.azure.core.util.Context.NONE);
    }
}
