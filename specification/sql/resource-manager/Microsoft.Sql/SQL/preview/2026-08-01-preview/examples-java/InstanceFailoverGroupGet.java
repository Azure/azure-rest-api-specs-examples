
/**
 * Samples for InstanceFailoverGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/InstanceFailoverGroupGet.json
     */
    /**
     * Sample code: Get failover group.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getFailoverGroup(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getInstanceFailoverGroups().getWithResponse("Default", "Japan East",
            "failover-group-test", com.azure.core.util.Context.NONE);
    }
}
