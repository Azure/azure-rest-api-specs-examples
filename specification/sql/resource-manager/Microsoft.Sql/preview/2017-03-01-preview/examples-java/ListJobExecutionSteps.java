import com.azure.core.util.Context;
import java.util.UUID;

/** Samples for JobStepExecutions ListByJobExecution. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/ListJobExecutionSteps.json
     */
    /**
     * Sample code: List job step executions.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listJobStepExecutions(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getJobStepExecutions()
            .listByJobExecution(
                "group1",
                "server1",
                "agent1",
                "job1",
                UUID.fromString("5555-6666-7777-8888-999999999999"),
                null,
                null,
                null,
                null,
                null,
                null,
                null,
                Context.NONE);
    }
}
