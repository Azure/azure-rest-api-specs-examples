
/**
 * Samples for FailoverGroups ForceFailoverAllowDataLoss.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/FailoverGroupForceFailoverAllowDataLoss.json
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
