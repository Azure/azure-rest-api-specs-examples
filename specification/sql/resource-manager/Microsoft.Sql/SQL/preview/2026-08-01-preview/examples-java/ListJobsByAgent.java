
/**
 * Samples for Jobs ListByAgent.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ListJobsByAgent.json
     */
    /**
     * Sample code: List jobs in a job agent.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listJobsInAJobAgent(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobs().listByAgent("group1", "server1", "agent1", com.azure.core.util.Context.NONE);
    }
}
