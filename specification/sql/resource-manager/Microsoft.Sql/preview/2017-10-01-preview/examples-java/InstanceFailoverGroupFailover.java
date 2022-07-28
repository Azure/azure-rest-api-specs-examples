import com.azure.core.util.Context;

/** Samples for InstanceFailoverGroups Failover. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-10-01-preview/examples/InstanceFailoverGroupFailover.json
     */
    /**
     * Sample code: Planned failover of a failover group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void plannedFailoverOfAFailoverGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getInstanceFailoverGroups()
            .failover("Default", "Japan West", "failover-group-test-3", Context.NONE);
    }
}
