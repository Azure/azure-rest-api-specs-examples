
/**
 * Samples for FailoverGroups ForceFailoverAllowDataLoss.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/FailoverGroupForceFailoverAllowDataLoss.json
     */
    /**
     * Sample code: Forced failover of a failover group allowing data loss.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        forcedFailoverOfAFailoverGroupAllowingDataLoss(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getFailoverGroups().forceFailoverAllowDataLoss("Default",
            "failover-group-secondary-server", "failover-group-test-3", com.azure.core.util.Context.NONE);
    }
}
