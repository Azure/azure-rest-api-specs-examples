
import java.util.UUID;

/**
 * Samples for JobTargetExecutions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/GetJobExecutionTarget.json
     */
    /**
     * Sample code: Get a job step target execution.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAJobStepTargetExecution(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getJobTargetExecutions().getWithResponse("group1", "server1",
            "agent1", "job1", UUID.fromString("5A86BF65-43AC-F258-2524-9E92992F97CA"), "step1",
            UUID.fromString("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"), com.azure.core.util.Context.NONE);
    }
}
