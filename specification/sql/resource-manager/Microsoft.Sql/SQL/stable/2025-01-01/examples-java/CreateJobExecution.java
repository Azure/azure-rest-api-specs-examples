
/**
 * Samples for JobExecutions Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/CreateJobExecution.json
     */
    /**
     * Sample code: Start a job execution.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void startAJobExecution(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobExecutions().create("group1", "server1", "agent1", "job1",
            com.azure.core.util.Context.NONE);
    }
}
