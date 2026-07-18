
/**
 * Samples for JobExecutions ListByAgent.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ListJobExecutionsByAgent.json
     */
    /**
     * Sample code: List all job executions in a job agent.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listAllJobExecutionsInAJobAgent(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobExecutions().listByAgent("group1", "server1", "agent1", null, null, null, null,
            null, null, null, com.azure.core.util.Context.NONE);
    }
}
