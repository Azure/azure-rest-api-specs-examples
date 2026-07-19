
/**
 * Samples for JobSteps ListByVersion.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ListJobStepsByVersion.json
     */
    /**
     * Sample code: List job steps for the specified version of a job.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        listJobStepsForTheSpecifiedVersionOfAJob(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobSteps().listByVersion("group1", "server1", "agent1", "job1", 1,
            com.azure.core.util.Context.NONE);
    }
}
