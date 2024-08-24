
import java.util.UUID;

/**
 * Samples for JobExecutions Cancel.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/CancelJobExecution.json
     */
    /**
     * Sample code: Cancel a job execution.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cancelAJobExecution(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getJobExecutions().cancelWithResponse("group1", "server1",
            "agent1", "job1", UUID.fromString("5A86BF65-43AC-F258-2524-9E92992F97CA"),
            com.azure.core.util.Context.NONE);
    }
}
