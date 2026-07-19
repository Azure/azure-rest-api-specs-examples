
import java.util.UUID;

/**
 * Samples for JobTargetExecutions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/GetJobExecutionTarget.json
     */
    /**
     * Sample code: Get a job step target execution.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getAJobStepTargetExecution(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobTargetExecutions().getWithResponse("group1", "server1", "agent1", "job1",
            UUID.fromString("5A86BF65-43AC-F258-2524-9E92992F97CA"), "step1",
            UUID.fromString("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"), com.azure.core.util.Context.NONE);
    }
}
