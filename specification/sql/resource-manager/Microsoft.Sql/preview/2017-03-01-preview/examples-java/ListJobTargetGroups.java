import com.azure.core.util.Context;

/** Samples for JobTargetGroups ListByAgent. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/ListJobTargetGroups.json
     */
    /**
     * Sample code: Get all target groups in an agent.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAllTargetGroupsInAnAgent(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getJobTargetGroups()
            .listByAgent("group1", "server1", "agent1", Context.NONE);
    }
}
