
import com.azure.core.util.Context;
import java.util.UUID;

/** Samples for JobExecutions CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/CreateOrUpdateJobExecution.json
     */
    /**
     * Sample code: Create job execution.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createJobExecution(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getJobExecutions().createOrUpdate("group1", "server1", "agent1",
            "job1", UUID.fromString("5A86BF65-43AC-F258-2524-9E92992F97CA"), Context.NONE);
    }
}
