
/**
 * Samples for JobAgents Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/GetJobAgent.json
     */
    /**
     * Sample code: Get a job agent.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getAJobAgent(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobAgents().getWithResponse("group1", "server1", "agent1",
            com.azure.core.util.Context.NONE);
    }
}
