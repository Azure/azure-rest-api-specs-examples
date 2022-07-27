import com.azure.core.util.Context;

/** Samples for Jobs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/GetJob.json
     */
    /**
     * Sample code: Get a job.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAJob(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getJobs()
            .getWithResponse("group1", "server1", "agent1", "job1", Context.NONE);
    }
}
