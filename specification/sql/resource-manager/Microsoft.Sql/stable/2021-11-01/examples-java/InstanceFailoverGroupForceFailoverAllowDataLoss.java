
/**
 * Samples for InstanceFailoverGroups ForceFailoverAllowDataLoss.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * InstanceFailoverGroupForceFailoverAllowDataLoss.json
     */
    /**
     * Sample code: Forced failover of a failover group allowing data loss.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        forcedFailoverOfAFailoverGroupAllowingDataLoss(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getInstanceFailoverGroups().forceFailoverAllowDataLoss("Default",
            "Japan West", "failover-group-test-3", com.azure.core.util.Context.NONE);
    }
}
