
/**
 * Samples for JobTargetGroups ListByAgent.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ListJobTargetGroups.json
     */
    /**
     * Sample code: Get all target groups in an agent.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAllTargetGroupsInAnAgent(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getJobTargetGroups().listByAgent("group1", "server1", "agent1",
            com.azure.core.util.Context.NONE);
    }
}
