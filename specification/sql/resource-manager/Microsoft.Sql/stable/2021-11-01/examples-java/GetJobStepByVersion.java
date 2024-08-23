
/**
 * Samples for JobSteps GetByVersion.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/GetJobStepByVersion.json
     */
    /**
     * Sample code: Get the specified version of a job step.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getTheSpecifiedVersionOfAJobStep(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getJobSteps().getByVersionWithResponse("group1", "server1",
            "agent1", "job1", 1, "step1", com.azure.core.util.Context.NONE);
    }
}
