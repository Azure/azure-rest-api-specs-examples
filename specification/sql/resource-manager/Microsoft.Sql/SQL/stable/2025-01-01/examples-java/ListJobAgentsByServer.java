
/**
 * Samples for JobAgents ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ListJobAgentsByServer.json
     */
    /**
     * Sample code: List job agents in a server.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listJobAgentsInAServer(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobAgents().listByServer("group1", "server1", com.azure.core.util.Context.NONE);
    }
}
