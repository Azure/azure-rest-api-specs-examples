
/**
 * Samples for JobVersions ListByJob.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ListJobVersions.json
     */
    /**
     * Sample code: Get all versions of a job.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAllVersionsOfAJob(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getJobVersions().listByJob("group1", "server1", "agent1", "job1",
            com.azure.core.util.Context.NONE);
    }
}
