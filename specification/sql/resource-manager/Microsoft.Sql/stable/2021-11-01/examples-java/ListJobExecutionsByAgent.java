
import com.azure.core.util.Context;

/** Samples for JobExecutions ListByAgent. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ListJobExecutionsByAgent.json
     */
    /**
     * Sample code: List all job executions in a job agent.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllJobExecutionsInAJobAgent(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getJobExecutions().listByAgent("group1", "server1", "agent1", null,
            null, null, null, null, null, null, Context.NONE);
    }
}
