
/**
 * Samples for JobSteps Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DeleteJobStep.json
     */
    /**
     * Sample code: Delete a job step.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deleteAJobStep(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobSteps().deleteWithResponse("group1", "server1", "agent1", "job1", "step1",
            com.azure.core.util.Context.NONE);
    }
}
