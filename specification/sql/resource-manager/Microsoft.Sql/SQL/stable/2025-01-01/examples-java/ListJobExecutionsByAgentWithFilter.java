
import java.time.OffsetDateTime;

/**
 * Samples for JobExecutions ListByAgent.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ListJobExecutionsByAgentWithFilter.json
     */
    /**
     * Sample code: List all job executions in a job agent with filtering.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        listAllJobExecutionsInAJobAgentWithFiltering(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobExecutions().listByAgent("group1", "server1", "agent1",
            OffsetDateTime.parse("2017-03-21T19:00:00Z"), OffsetDateTime.parse("2017-03-21T19:05:00Z"),
            OffsetDateTime.parse("2017-03-21T19:20:00Z"), OffsetDateTime.parse("2017-03-21T19:25:00Z"), false, null,
            null, com.azure.core.util.Context.NONE);
    }
}
