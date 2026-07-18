
import java.util.UUID;

/**
 * Samples for JobTargetExecutions ListByJobExecution.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ListJobExecutionTargetsByExecution.json
     */
    /**
     * Sample code: List job step target executions.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listJobStepTargetExecutions(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobTargetExecutions().listByJobExecution("group1", "server1", "agent1", "job1",
            UUID.fromString("5A86BF65-43AC-F258-2524-9E92992F97CA"), null, null, null, null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
