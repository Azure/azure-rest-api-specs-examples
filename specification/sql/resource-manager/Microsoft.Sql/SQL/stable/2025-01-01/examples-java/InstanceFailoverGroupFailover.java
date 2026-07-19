
/**
 * Samples for InstanceFailoverGroups Failover.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/InstanceFailoverGroupFailover.json
     */
    /**
     * Sample code: Planned failover of a failover group.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void plannedFailoverOfAFailoverGroup(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getInstanceFailoverGroups().failover("Default", "Japan West", "failover-group-test-3",
            com.azure.core.util.Context.NONE);
    }
}
