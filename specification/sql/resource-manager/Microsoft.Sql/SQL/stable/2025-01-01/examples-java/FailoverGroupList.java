
/**
 * Samples for FailoverGroups ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/FailoverGroupList.json
     */
    /**
     * Sample code: List failover group.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listFailoverGroup(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getFailoverGroups().listByServer("Default", "failover-group-primary-server",
            com.azure.core.util.Context.NONE);
    }
}
