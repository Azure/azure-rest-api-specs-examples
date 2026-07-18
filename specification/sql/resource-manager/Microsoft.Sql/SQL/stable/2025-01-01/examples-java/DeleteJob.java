
/**
 * Samples for Jobs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DeleteJob.json
     */
    /**
     * Sample code: Delete a job.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deleteAJob(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobs().deleteWithResponse("group1", "server1", "agent1", "job1",
            com.azure.core.util.Context.NONE);
    }
}
