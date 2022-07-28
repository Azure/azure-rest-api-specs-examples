import com.azure.core.util.Context;

/** Samples for JobVersions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/GetJobVersion.json
     */
    /**
     * Sample code: Get a version of a job.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAVersionOfAJob(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getJobVersions()
            .getWithResponse("group1", "server1", "agent1", "job1", 1, Context.NONE);
    }
}
