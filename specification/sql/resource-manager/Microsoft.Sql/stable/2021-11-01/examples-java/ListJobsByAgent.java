
import com.azure.core.util.Context;

/** Samples for Jobs ListByAgent. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ListJobsByAgent.json
     */
    /**
     * Sample code: List jobs in a job agent.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listJobsInAJobAgent(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getJobs().listByAgent("group1", "server1", "agent1", Context.NONE);
    }
}
