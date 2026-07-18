
/**
 * Samples for JobSteps Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/GetJobStepByJob.json
     */
    /**
     * Sample code: Get the latest version of a job step.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getTheLatestVersionOfAJobStep(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobSteps().getWithResponse("group1", "server1", "agent1", "job1", "step1",
            com.azure.core.util.Context.NONE);
    }
}
