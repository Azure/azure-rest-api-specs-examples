import com.azure.core.util.Context;

/** Samples for JobTargetGroups Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/GetJobTargetGroup.json
     */
    /**
     * Sample code: Get a target group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getATargetGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getJobTargetGroups()
            .getWithResponse("group1", "server1", "agent1", "targetGroup1", Context.NONE);
    }
}
