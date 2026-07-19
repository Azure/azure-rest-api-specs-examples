
/**
 * Samples for JobVersions ListByJob.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ListJobVersions.json
     */
    /**
     * Sample code: Get all versions of a job.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getAllVersionsOfAJob(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobVersions().listByJob("group1", "server1", "agent1", "job1",
            com.azure.core.util.Context.NONE);
    }
}
