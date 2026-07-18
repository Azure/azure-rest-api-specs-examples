
/**
 * Samples for JobExecutions ListByJob.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ListJobExecutionsByJob.json
     */
    /**
     * Sample code: List a job's executions.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listAJobSExecutions(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobExecutions().listByJob("group1", "server1", "agent1", "job1", null, null, null,
            null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
