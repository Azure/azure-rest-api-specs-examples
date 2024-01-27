
import com.azure.core.util.Context;

/** Samples for JobExecutions ListByJob. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ListJobExecutionsByJob.json
     */
    /**
     * Sample code: List a job's executions.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAJobSExecutions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getJobExecutions().listByJob("group1", "server1", "agent1", "job1",
            null, null, null, null, null, null, null, Context.NONE);
    }
}
