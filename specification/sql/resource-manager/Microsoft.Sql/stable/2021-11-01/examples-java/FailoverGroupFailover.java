
/**
 * Samples for FailoverGroups Failover.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/FailoverGroupFailover.json
     */
    /**
     * Sample code: Planned failover of a failover group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void plannedFailoverOfAFailoverGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getFailoverGroups().failover("Default",
            "failover-group-secondary-server", "failover-group-test-3", com.azure.core.util.Context.NONE);
    }
}
