import com.azure.core.util.Context;

/** Samples for JobVersions ListByJob. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/ListJobVersions.json
     */
    /**
     * Sample code: Get all versions of a job.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAllVersionsOfAJob(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getJobVersions()
            .listByJob("group1", "server1", "agent1", "job1", Context.NONE);
    }
}
