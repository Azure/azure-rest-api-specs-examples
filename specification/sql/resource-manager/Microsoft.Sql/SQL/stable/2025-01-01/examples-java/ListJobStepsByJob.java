
/**
 * Samples for JobSteps ListByJob.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ListJobStepsByJob.json
     */
    /**
     * Sample code: List job steps for the latest version of a job.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listJobStepsForTheLatestVersionOfAJob(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobSteps().listByJob("group1", "server1", "agent1", "job1",
            com.azure.core.util.Context.NONE);
    }
}
