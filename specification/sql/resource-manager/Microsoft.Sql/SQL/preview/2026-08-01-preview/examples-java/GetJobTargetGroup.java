
/**
 * Samples for JobTargetGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/GetJobTargetGroup.json
     */
    /**
     * Sample code: Get a target group.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getATargetGroup(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getJobTargetGroups().getWithResponse("group1", "server1", "agent1", "targetGroup1",
            com.azure.core.util.Context.NONE);
    }
}
