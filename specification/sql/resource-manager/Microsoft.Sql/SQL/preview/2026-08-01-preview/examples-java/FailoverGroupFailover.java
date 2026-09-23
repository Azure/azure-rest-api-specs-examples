
/**
 * Samples for FailoverGroups Failover.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/FailoverGroupFailover.json
     */
    /**
     * Sample code: Planned failover of a failover group.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void plannedFailoverOfAFailoverGroup(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getFailoverGroups().failover("Default", "failover-group-secondary-server",
            "failover-group-test-3", com.azure.core.util.Context.NONE);
    }
}
