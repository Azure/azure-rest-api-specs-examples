
import java.util.UUID;

/**
 * Samples for JobStepExecutions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/GetJobExecutionStep.json
     */
    /**
     * Sample code: Get a job step execution.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getAJobStepExecution(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobStepExecutions().getWithResponse("group1", "server1", "agent1", "job1",
            UUID.fromString("5A86BF65-43AC-F258-2524-9E92992F97CA"), "step1", com.azure.core.util.Context.NONE);
    }
}
