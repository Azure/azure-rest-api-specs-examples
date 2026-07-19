
import java.util.UUID;

/**
 * Samples for JobExecutions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/CreateOrUpdateJobExecution.json
     */
    /**
     * Sample code: Create job execution.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void createJobExecution(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobExecutions().createOrUpdate("group1", "server1", "agent1", "job1",
            UUID.fromString("5A86BF65-43AC-F258-2524-9E92992F97CA"), com.azure.core.util.Context.NONE);
    }
}
