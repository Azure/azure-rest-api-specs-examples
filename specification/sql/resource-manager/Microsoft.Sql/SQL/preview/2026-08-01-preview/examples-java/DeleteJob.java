
/**
 * Samples for Jobs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/DeleteJob.json
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
