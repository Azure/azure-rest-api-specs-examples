
import java.util.UUID;

/**
 * Samples for JobTargetExecutions ListByJobExecution.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ListJobExecutionTargetsByExecution.
     * json
     */
    /**
     * Sample code: List job step target executions.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listJobStepTargetExecutions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getJobTargetExecutions().listByJobExecution("group1", "server1",
            "agent1", "job1", UUID.fromString("5A86BF65-43AC-F258-2524-9E92992F97CA"), null, null, null, null, null,
            null, null, com.azure.core.util.Context.NONE);
    }
}
