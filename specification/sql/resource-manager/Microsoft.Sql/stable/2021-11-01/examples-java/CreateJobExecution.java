
/**
 * Samples for JobExecutions Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/CreateJobExecution.json
     */
    /**
     * Sample code: Start a job execution.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void startAJobExecution(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getJobExecutions().create("group1", "server1", "agent1", "job1",
            com.azure.core.util.Context.NONE);
    }
}
