
/**
 * Samples for FailoverGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/FailoverGroupGet.json
     */
    /**
     * Sample code: Get failover group.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getFailoverGroup(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getFailoverGroups().getWithResponse("Default", "failovergroupprimaryserver",
            "failovergrouptest3", com.azure.core.util.Context.NONE);
    }
}
