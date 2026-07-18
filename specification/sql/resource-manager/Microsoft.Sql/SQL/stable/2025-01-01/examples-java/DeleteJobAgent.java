
/**
 * Samples for JobAgents Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DeleteJobAgent.json
     */
    /**
     * Sample code: Delete a job agent.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deleteAJobAgent(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobAgents().delete("group1", "server1", "agent1", com.azure.core.util.Context.NONE);
    }
}
