
/**
 * Samples for JobSteps GetByVersion.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/GetJobStepByVersion.json
     */
    /**
     * Sample code: Get the specified version of a job step.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getTheSpecifiedVersionOfAJobStep(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobSteps().getByVersionWithResponse("group1", "server1", "agent1", "job1", 1,
            "step1", com.azure.core.util.Context.NONE);
    }
}
