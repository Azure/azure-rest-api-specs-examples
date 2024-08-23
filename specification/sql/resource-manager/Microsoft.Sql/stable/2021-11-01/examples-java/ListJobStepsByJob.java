
/**
 * Samples for JobSteps ListByJob.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ListJobStepsByJob.json
     */
    /**
     * Sample code: List job steps for the latest version of a job.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listJobStepsForTheLatestVersionOfAJob(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getJobSteps().listByJob("group1", "server1", "agent1", "job1",
            com.azure.core.util.Context.NONE);
    }
}
