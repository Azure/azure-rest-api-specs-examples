
/**
 * Samples for InstanceFailoverGroups ForceFailoverAllowDataLoss.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/InstanceFailoverGroupForceFailoverAllowDataLoss.json
     */
    /**
     * Sample code: Forced failover of a failover group allowing data loss.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        forcedFailoverOfAFailoverGroupAllowingDataLoss(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getInstanceFailoverGroups().forceFailoverAllowDataLoss("Default", "Japan West",
            "failover-group-test-3", com.azure.core.util.Context.NONE);
    }
}
