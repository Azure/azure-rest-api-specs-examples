
/**
 * Samples for JobSteps Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DeleteJobStep.json
     */
    /**
     * Sample code: Delete a job step.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAJobStep(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getJobSteps().deleteWithResponse("group1", "server1", "agent1",
            "job1", "step1", com.azure.core.util.Context.NONE);
    }
}
