import com.azure.core.util.Context;
import java.util.UUID;

/** Samples for JobExecutions CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/CreateOrUpdateJobExecution.json
     */
    /**
     * Sample code: Create job execution.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createJobExecution(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getJobExecutions()
            .createOrUpdate(
                "group1",
                "server1",
                "agent1",
                "job1",
                UUID.fromString("5555-6666-7777-8888-999999999999"),
                Context.NONE);
    }
}
