
/**
 * Samples for Jobs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/GetJob.json
     */
    /**
     * Sample code: Get a job.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getAJob(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobs().getWithResponse("group1", "server1", "agent1", "job1",
            com.azure.core.util.Context.NONE);
    }
}
