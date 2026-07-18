
/**
 * Samples for JobTargetGroups ListByAgent.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ListJobTargetGroups.json
     */
    /**
     * Sample code: Get all target groups in an agent.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getAllTargetGroupsInAnAgent(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobTargetGroups().listByAgent("group1", "server1", "agent1",
            com.azure.core.util.Context.NONE);
    }
}
