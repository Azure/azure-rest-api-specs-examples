
import com.azure.resourcemanager.sql.fluent.models.JobInner;

/**
 * Samples for Jobs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/CreateOrUpdateJobMin.json
     */
    /**
     * Sample code: Create a job with default properties.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAJobWithDefaultProperties(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getJobs().createOrUpdateWithResponse("group1", "server1", "agent1",
            "job1", new JobInner(), com.azure.core.util.Context.NONE);
    }
}
