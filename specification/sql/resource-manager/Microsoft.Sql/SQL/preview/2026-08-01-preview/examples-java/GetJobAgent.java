
/**
 * Samples for JobAgents Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/GetJobAgent.json
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
