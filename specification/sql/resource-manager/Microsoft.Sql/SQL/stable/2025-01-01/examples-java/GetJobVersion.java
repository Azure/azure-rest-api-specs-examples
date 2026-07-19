
/**
 * Samples for JobVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/GetJobVersion.json
     */
    /**
     * Sample code: Get a version of a job.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getAVersionOfAJob(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobVersions().getWithResponse("group1", "server1", "agent1", "job1", 1,
            com.azure.core.util.Context.NONE);
    }
}
