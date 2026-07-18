
/**
 * Samples for JobTargetGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DeleteJobTargetGroup.json
     */
    /**
     * Sample code: Delete a target group.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deleteATargetGroup(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobTargetGroups().deleteWithResponse("group1", "server1", "agent1", "targetGroup1",
            com.azure.core.util.Context.NONE);
    }
}
