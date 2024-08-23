
/**
 * Samples for JobTargetGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DeleteJobTargetGroup.json
     */
    /**
     * Sample code: Delete a target group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteATargetGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getJobTargetGroups().deleteWithResponse("group1", "server1",
            "agent1", "targetGroup1", com.azure.core.util.Context.NONE);
    }
}
