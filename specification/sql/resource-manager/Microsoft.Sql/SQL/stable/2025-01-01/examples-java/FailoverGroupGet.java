
/**
 * Samples for FailoverGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/FailoverGroupGet.json
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
